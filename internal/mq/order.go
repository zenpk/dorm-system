package mq

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/zap"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"os/signal"
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

func (o *OrderProducer) Send(req *order.SubmitRequest) error {
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
}

// Init consumer and subscribe
func (o *OrderConsumer) Init(c *viper.Viper) error {
	o.config = c
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
			log.Fatalf("failed to close consumer, error: %v", err)
		}
	}()
	partitionList, err := o.consumer.Partitions(o.config.GetString("kafka.topic"))
	if err != nil {
		return err
	}
	for _, partition := range partitionList {
		partitionConsumer, err := o.consumer.ConsumePartition(o.config.GetString("kafka.topic"), partition, sarama.OffsetOldest)
		if err != nil {
			return err
		}
		// Trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		// start a go routine for every partitionConsumer
		go func() {
			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalf("failed to close partitionConsumer, error: %v", err)
				}
			}()
		ConsumerLoop:
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					// start a go routine to handle order
					go func() {
						ctx, cancel := util.ContextWithTimeout(o.config.GetInt("timeout.submit"))
						defer cancel()
						zap.Logger.Infof("consumed message offset %d\n", msg.Offset)
						if err := order.Submit(ctx, msg); err != nil {
							zap.Logger.Error(err)
						}
					}()
				case <-signals:
					break ConsumerLoop
				}
			}
		}()
	}
	return nil
}
