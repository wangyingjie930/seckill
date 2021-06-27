package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

const MQURL = "amqp://admin:admin@192.168.205.10:5672/"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel

	//会用到的就是这三种参数的组合
	Exchange string
	Queue string
	Key string

	Mqurl string
}

func NewRabbitMQ(queue, exchange, key string) *RabbitMQ {
	r := &RabbitMQ{Exchange: exchange, Queue: queue, Key: key, Mqurl: MQURL}
	var err error
	r.conn, err = amqp.Dial(r.Mqurl)
	r.FailOnError(err, "connect error")
	r.channel, err = r.conn.Channel()
	r.FailOnError(err, "channel error")
	return r
}

func (r *RabbitMQ) Destroy() {
	err := r.conn.Close()
	r.FailOnError(err, "connect close error")
	err = r.channel.Close()
	r.FailOnError(err, "channel close error")
}

func (r *RabbitMQ) FailOnError (err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}