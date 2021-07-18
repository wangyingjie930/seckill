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
		//自动删除的前提是: 至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会 自动删除
		false,
		//如果一个队列被声明为排 他队列，该队列仅对首次声明它的连接可见，并在连接断开时自动删除。
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
		false,  // 如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,  // 如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
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
		//用来区分多个消费者
		"",     // consumer
		//是否自动应答
		true,   // auto-ack
		//是否独有
		false,  // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,  // no-local
		//列是否阻塞
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
