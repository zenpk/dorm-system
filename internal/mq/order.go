package mq

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/pkg/zap"
	"log"
	"os"
	"os/signal"
	"time"
)

type OrderProducer struct {
	Producer sarama.AsyncProducer
}

// Init create topic and initialize producer
func (o *OrderProducer) Init() error {
	// create topic
	detail := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	if err := ClusterAdmin.CreateTopic(viper.GetString("kafka.topics.order"), detail, false); err != nil {
		return err
	}
	return o.initProducer()
}

func (o *OrderProducer) initProducer() error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	var err error
	o.Producer, err = sarama.NewAsyncProducer(viper.GetStringSlice("kafka.brokers"), config)
	return err
}

func (o *OrderProducer) Send(req *dto.OrderRequest) error {
	reqByte, err := json.Marshal(req)
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: viper.GetString("kafka.order.topic"),
		Value: sarama.ByteEncoder(reqByte),
	}
	o.Producer.Input() <- msg
	return nil
}

type OrderConsumer struct {
	Consumer sarama.Consumer
}

// Init consumer and subscribe
func (o *OrderConsumer) Init() error {
	// init consumer
	config := sarama.NewConfig()
	var err error
	o.Consumer, err = sarama.NewConsumer(viper.GetStringSlice("kafka.brokers"), config)
	if err != nil {
		return err
	}
	return o.subscribe()
}

func (o *OrderConsumer) subscribe() error {
	defer func() {
		if err := o.Consumer.Close(); err != nil {
			log.Fatalf("failed to close consumer: %v", err)
		}
	}()
	partitionList, err := o.Consumer.Partitions(viper.GetString("kafka.topics.order"))
	if err != nil {
		return err
	}
	for _, partition := range partitionList {
		partitionConsumer, err := o.Consumer.ConsumePartition(viper.GetString("kafka.topics.order"), partition, sarama.OffsetOldest)
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
