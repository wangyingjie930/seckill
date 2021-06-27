package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ(queueName, "", "")
	return rabbitMQ
}

func (r *RabbitMQ) PublishSimple(message string) {
	_, err := r.channel.QueueDeclare(
		r.Queue,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性, 其他账号不可访问
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	r.FailOnError(err, "Failed to declare a queue")

	err = r.channel.Publish(
		r.Exchange,     // exchange
		r.Queue, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.FailOnError(err, "Failed to publish a message")
	fmt.Print("send message!")
}

func (r *RabbitMQ) ConsumeSimple()  {
	_, err := r.channel.QueueDeclare(
		r.Queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	r.FailOnError(err, "Failed to declare a queue")

	msgs, err := r.channel.Consume(
		r.Queue, // queue
		"",     // consumer
		true,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	r.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever
}
