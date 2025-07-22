package mq

type Producer struct{}

func NewProducer() *Producer {
	// 初始化 MQ 生产者
	return &Producer{}
}
