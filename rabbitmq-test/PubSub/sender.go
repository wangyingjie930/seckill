package main

import (
	"fmt"
	"imooc-product/rabbitmq-test/rabbitmq"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rabbitMQ := rabbitmq.NewRabbitMQPubSub("publish/subscribe mode")
	for i := 0; i <= 100; i++ {
		rabbitMQ.PublishPub("hello" + strconv.Itoa(rand.Int()))
		time.Sleep(3 * time.Second)
		fmt.Println(i)
	}
}
