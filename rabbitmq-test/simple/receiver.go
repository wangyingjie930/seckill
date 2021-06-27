package main

import "imooc-product/rabbitmq-test/rabbitmq"

func main() {
	rabbitMQ := rabbitmq.NewRabbitMQSimple("simple")
	rabbitMQ.ConsumeSimple()
}
