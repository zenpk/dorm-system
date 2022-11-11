package mq

func InitProducer() error {
	if err := Producer.Order.InitProducer(); err != nil {
		return err
	}
	return nil
}
