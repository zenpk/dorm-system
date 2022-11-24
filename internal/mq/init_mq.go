package mq

import (
	"github.com/Shopify/sarama"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
)

var ClusterAdmin sarama.ClusterAdmin

// InitMQ initialize cluster admin and all producers, return all producers
func InitMQ() ([]sarama.AsyncProducer, error) {
	// initialize cluster
	// read all configs
	brokers := make([]string, 0)
	allProducers := make([]sarama.AsyncProducer, 0)
	orderConfig, err := viperpkg.InitConfig("order")
	if err != nil {
		return nil, err
	}
	brokers = append(brokers, orderConfig.GetStringSlice("kafka.brokers")...)
	config := sarama.NewConfig()
	ClusterAdmin, err = sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return nil, err
	}

	// initialize producers
	// order
	orderProducer, err := Producer.Order.init(orderConfig)
	if err != nil {
		return nil, err
	}
	allProducers = append(allProducers, orderProducer)
	return allProducers, nil
}
