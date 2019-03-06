package main

import (
	"log"

	"github.com/streadway/amqp"
)

func (rmq *RMQ) Receive(queue string) {
	// open connection
	conn, err := amqp.Dial(rmq.url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() {
		log.Panicln("Closing connection")
		cErr := conn.Close()
		failOnError(cErr, "Failed to close RMQ connection")
	}()

	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() {
		log.Panicln("Closing channel")
		cErr := ch.Close()
		failOnError(cErr, "Failed to close RMQ channel")
	}()

	q, err := ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			body := string(d.Body)
			switch body {
			case "panic":
				log.Panic("PANIC!")
			default:
				log.Printf("Received a message: %s", body)
			}
			d.Ack(true)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
