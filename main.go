package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	checkInputArguments(args)

	cmd := args[0]
	const queueName = "RMQ-HELLO-RQ"
	if cmd == "s" {
		Send(queueName, args[1])
	} else if cmd == "r" {
		Receive(queueName)
	} else {
		log.Println("Unknown command")
	}

}

func checkInputArguments(args []string) {
	if len(args) == 0 {
		log.Fatalln("Supply arguments please")
	}
}
