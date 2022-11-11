package mq

type ProducerSet struct {
	Order OrderProducer
}

type ConsumerSet struct {
	Order OrderConsumer
}

var Producer ProducerSet
var Consumer ConsumerSet
