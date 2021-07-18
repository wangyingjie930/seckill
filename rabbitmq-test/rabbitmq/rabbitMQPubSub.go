package rabbitmq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

func (r *RabbitMQ) PublishPub (message string) {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	r.FailOnError(err, "exchange declare error")

	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain", Body: []byte(message),
		})
	r.FailOnError(err, "Publish error")
	fmt.Print("publish data!\n")
}

func (r *RabbitMQ) ConsumePub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil)
	r.FailOnError(err, "exchange declare error")

	q, err := r.channel.QueueDeclare(
		r.Queue,
		false,
		false,
		false,
		false,
		nil)
	r.FailOnError(err, "queue declare error")
	log.Printf("queue name %s", q.Name)

	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil)
	r.FailOnError(err, "queue bind error")

	msgs, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)
	for msg := range msgs {
		log.Printf(" [x] %s\n", msg.Body)
	}
	<-forever
}