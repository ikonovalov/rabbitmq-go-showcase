package main

import (
	"io"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func CloseResource(c io.Closer) {
	cErr := c.Close()
	failOnError(cErr, "Failed to close RMQ connection")
}
