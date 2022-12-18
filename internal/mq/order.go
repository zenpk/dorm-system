package mq

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/zap"
	"google.golang.org/protobuf/proto"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

type OrderProducer struct {
	config   *viper.Viper
	producer sarama.AsyncProducer
}

// init create topic and producer
func (o *OrderProducer) init(c *viper.Viper) (sarama.AsyncProducer, error) {
	o.config = c
	// check if topic exists
	topicMap, err := ClusterAdmin.ListTopics()
	if err != nil {
		return nil, err
	}
	topic := o.config.GetString("kafka.topic")
	if _, ok := topicMap[topic]; !ok { // topic doesn't exist
		// create topic
		detail := &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}
		if err := ClusterAdmin.CreateTopic(topic, detail, false); err != nil {
			return nil, err
		}
	}
	// create producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	o.producer, err = sarama.NewAsyncProducer(o.config.GetStringSlice("kafka.brokers"), config)
	return o.producer, err
}

func (o *OrderProducer) Send(req *pb.SubmitRequest) error {
	reqByte, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: o.config.GetString("kafka.topic"),
		Value: sarama.ByteEncoder(reqByte),
	}
	o.producer.Input() <- msg
	return nil
}

type OrderConsumer struct {
	config   *viper.Viper
	consumer sarama.Consumer
	server   *pb.Server
}

// Init consumer and subscribe
func (o *OrderConsumer) Init(server *pb.Server) error {
	o.config = server.Config
	o.server = server
	// init consumer
	config := sarama.NewConfig()
	var err error
	o.consumer, err = sarama.NewConsumer(o.config.GetStringSlice("kafka.brokers"), config)
	if err != nil {
		return err
	}
	return o.subscribe()
}

// subscribe to message queue to handle submitted orders
func (o *OrderConsumer) subscribe() error {
	defer func() {
		if err := o.consumer.Close(); err != nil {
			log.Fatalf("failed to close consumer: %v", err)
		}
	}()
	partitionList, err := o.consumer.Partitions(o.config.GetString("kafka.topic"))
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	for _, partition := range partitionList {
		partitionConsumer, err := o.consumer.ConsumePartition(o.config.GetString("kafka.topic"), partition, sarama.OffsetOldest)
		if err != nil {
			return err
		}
		// Trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		// start a go routine for every partitionConsumer
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer func() {
				defer wg.Done()
				if err := pc.Close(); err != nil {
					log.Fatalf("failed to close partitionConsumer: %v", err)
				}
			}()
		ConsumerLoop:
			for {
				select {
				case msg := <-pc.Messages():
					var req pb.SubmitRequest
					if err := proto.Unmarshal(msg.Value, &req); err != nil {
						zap.Logger.Error(err)
						continue
					}
					timeout := o.config.GetInt("timeout.submit")
					ctx, cancel := util.ContextWithTimeout(timeout)
					zap.Logger.Infof("consumed message offset %d\n", msg.Offset)
					if _, err := o.server.Submit(ctx, &req); err != nil {
						zap.Logger.Errorf("consume message failed, offset: %v: %v", msg.Offset, err)
						// wait random time to retry
						waitTime := time.Duration(rand.Intn(timeout)+1) * time.Second
						time.Sleep(waitTime)
						if _, err := o.server.Submit(ctx, &req); err != nil {
							zap.Logger.Errorf("consume message failed for the 2nd time, offset: %v: %v", msg.Offset, err)
							// still failed, mark the order as failed in database
							if err := o.server.Fail(ctx, &req, "retried too many times"); err != nil {
								zap.Logger.Errorf("message completely failed! offset: %v: %v", msg.Offset, err)
							}
						}
					}
					cancel()
				case <-signals:
					break ConsumerLoop
				}
			}
		}(partitionConsumer)
	}
	wg.Wait()
	return nil
}
