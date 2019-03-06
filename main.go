package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	CheckInputArguments(args)

	cmd := args[0]
	const queueName = "RMQ-HELLO-RQ"
	switch cmd {
	case "s":
		payload := Payload(args)
		times := RepeatTime(args)

		Send(queueName, payload, times)
	case "r":
		Receive(queueName)
	default:
		log.Println("Unknown command")
	}
}

func Payload(args []string) string {
	if len(args) > 1 && args[1] != "_" {
		return args[1]
	} else {
		return NowTime()
	}
}

func NowTime() string {
	return time.Now().Format(time.RFC3339)
}

func RepeatTime(args []string) (times int) {
	if len(args) > 2 {
		r, e := strconv.ParseInt(args[2], 10, 64)
		failOnError(e, "Wrong times number")
		times = int(r)
	} else {
		times = 1
	}
	return times
}

func CheckInputArguments(args []string) {
	if len(args) == 0 {
		log.Fatalln("Supply arguments please")
	}
}
