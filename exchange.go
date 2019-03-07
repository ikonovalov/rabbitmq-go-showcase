package main

import (
	"github.com/streadway/amqp"
	"log"
)

func (rmq *RMQ) SendExchange(exchange string, message string, times int) {
	conn, err := amqp.Dial(rmq.url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() { failOnError(conn.Close(), "Failed to close RMQ connection") }()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() { failOnError(ch.Close(), "Fail to close channel") }()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	body := message

	for i := 0; i < times; i++ {
		err = ch.Publish(
			exchange, // exchange
			"",       // routing key
			false,    // mandatory
			false,    // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Transient,
				ContentType:  "text/plain",
				Body:         []byte(body),
			})
		failOnError(err, "Failed to publish a message")
	}
	log.Println("Done")

}
