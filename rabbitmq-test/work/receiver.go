package main

import "imooc-product/rabbitmq-test/rabbitmq"

func main() {
	rabbitMQ := rabbitmq.NewRabbitMQSimple("work model")
	rabbitMQ.ConsumeSimple()
}
