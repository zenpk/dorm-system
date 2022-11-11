package mq

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/pkg/zap"
)

type OrderProducer struct {
	Producer sarama.SyncProducer
}

func (o *OrderProducer) InitProducer() error {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	var err error
	o.Producer, err = sarama.NewSyncProducer(viper.GetStringSlice("kafka.brokers"), config)
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
	partition, offset, err := o.Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	zap.Logger.Infof("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
	return nil
}

type OrderConsumer struct {
	Config   *viper.Viper
	Consumer sarama.Consumer
}

func (o *OrderConsumer) InitConsumer() error {
	config := sarama.NewConfig()
	var err error
	o.Consumer, err = sarama.NewConsumer(o.Config.GetStringSlice("kafka.brokers"), config)
	return err
}

func (o *OrderConsumer) Subscribe() error {
	partitionList, err := o.Consumer.Partitions(o.Config.GetString("kafka.topic"))
	if err != nil {
		return err
	}
	offset := sarama.OffsetOldest // get offset for the oldest message on the topic
	for _, partition := range partitionList {
		pc, err := o.Consumer.ConsumePartition(o.Config.GetString("kafka.topic"), partition, offset)
		if err != nil {
			return err
		}
		go func(pc sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-pc.Messages():
					if err := order.Submit(msg); err != nil {
						zap.Logger.Error(err)
						// TODO
					}
				}
			}
		}(pc)
	}
	for {
	}
	return nil
}
