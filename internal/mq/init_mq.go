package mq

import (
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

var ClusterAdmin sarama.ClusterAdmin

func InitMQ() error {
	// Init ClusterAdmin
	config := sarama.NewConfig()
	var err error
	ClusterAdmin, err = sarama.NewClusterAdmin(viper.GetStringSlice("kafka.broker"), config)
	if err != nil {
		return err
	}
	return InitProducer()
}

// InitProducer initialize all producers
func InitProducer() error {
	if err := Producer.Order.Init(); err != nil {
		return err
	}
	return nil
}
