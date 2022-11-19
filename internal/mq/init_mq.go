package mq

import (
	"github.com/Shopify/sarama"
	"github.com/zenpk/dorm-system/pkg/viperpkg"
)

var ClusterAdmin sarama.ClusterAdmin

func InitMQ() error {
	// initialize cluster
	// read all configs
	brokers := make([]string, 0)
	orderConfig, err := viperpkg.InitConfig("order")
	if err != nil {
		return err
	}
	brokers = append(brokers, orderConfig.GetStringSlice("kafka.brokers")...)
	config := sarama.NewConfig()
	ClusterAdmin, err = sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		return err
	}

	// initialize producers
	// order
	if err := Producer.Order.init(orderConfig); err != nil {
		return err
	}
	return nil
}
