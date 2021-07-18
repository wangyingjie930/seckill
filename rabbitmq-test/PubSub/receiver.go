package main

import "imooc-product/rabbitmq-test/rabbitmq"

func main() {
	rabbitMQ := rabbitmq.NewRabbitMQPubSub("publish/subscribe mode")
	rabbitMQ.ConsumePub()
}
