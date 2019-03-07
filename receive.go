package main

import (
	"log"

	"github.com/streadway/amqp"
)

func (rmq *RMQ) Receive(exchange string) {
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
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	err = ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	failOnError(err, "Fail to bind queue to exchange")

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
			ackError := d.Ack(false)
			failOnError(ackError, "ACK failed")

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
