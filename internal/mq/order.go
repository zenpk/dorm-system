package mq

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/service/order"
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
func (o *OrderProducer) init(c *viper.Viper) error {
	o.config = c
	// check if topic exists
	topicMap, err := ClusterAdmin.ListTopics()
	if err != nil {
		return err
	}
	topic := o.config.GetString("kafka.topic")
	if _, ok := topicMap[topic]; !ok { // topic doesn't exist
		// create topic
		detail := &sarama.TopicDetail{
			NumPartitions:     1,
			ReplicationFactor: 1,
		}
		if err := ClusterAdmin.CreateTopic(topic, detail, false); err != nil {
			return err
		}
	}
	// create producer
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	o.producer, err = sarama.NewAsyncProducer(o.config.GetStringSlice("kafka.brokers"), config)
	return err
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
	for _, partition := range partitionList {
		partitionConsumer, err := o.consumer.ConsumePartition(o.config.GetString("kafka.topic"), partition, sarama.OffsetOldest)
		if err != nil {
			return err
		}
		defer func() {
			if err := partitionConsumer.Close(); err != nil {
				log.Fatalf("failed to close partitionConsumer: %v", err)
			}
		}()
		// Trap SIGINT to trigger a shutdown.
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
	ConsumerLoop:
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				zap.Logger.Infof("consumed message offset %d\n", msg.Offset)
				if err := order.Submit(msg); err != nil {
					zap.Logger.Error(err)
				}
			case <-signals:
				break ConsumerLoop
			}
		}
	}
	return nil
}
