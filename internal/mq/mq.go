package mq

type AllProducer struct {
	Order OrderProducer
}

type AllConsumer struct {
	Order OrderConsumer
}

var Producer AllProducer
var Consumer AllConsumer
