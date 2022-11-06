package mq

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"time"
)

var Producer sarama.AsyncProducer
var Consumer sarama.Consumer

func InitProducer() (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	Producer, err = sarama.NewAsyncProducer(viper.GetStringSlice("kafka.brokers"), config)
	return
}

func InitConsumer() (err error) {
	config := sarama.NewConfig()
	Consumer, err = sarama.NewConsumer(viper.GetStringSlice("kafka.brokers"), config)
	return
}
