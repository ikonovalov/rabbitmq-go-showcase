package main

import (
	"github.com/streadway/amqp"
	"log"
)

func (rmq *RMQ) SendOne(queue string, message string) {
	rmq.Send(queue, message, 1)
}

func (rmq *RMQ) Send(queue string, message string, times int) {
	conn, err := amqp.Dial(rmq.url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		cErr := conn.Close()
		failOnError(cErr, "Failed to close RMQ connection")
	}()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := message

	for i := 0; i < times; i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}
	log.Println("Done")

}
