package main

import (
	"imooc-product/rabbitmq-test/rabbitmq"
	"time"
)

func main() {
	rabbitMQ := rabbitmq.NewRabbitMQSimple("simple")
	for true {
		rabbitMQ.PublishSimple("hello")
		time.Sleep(3 * time.Second)
	}
}
