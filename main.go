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

		var i int64 = 0
		for ; i < times; i++ {
			Send(queueName, payload)
		}
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
	t := time.Now()
	return t.Format(time.RFC3339)
}

func RepeatTime(args []string) int64 {
	if len(args) > 2 {
		r, e := strconv.ParseInt(args[2], 10, 64)
		failOnError(e, "Wrong times number")
		return r
	} else {
		return 1
	}
}

func CheckInputArguments(args []string) {
	if len(args) == 0 {
		log.Fatalln("Supply arguments please")
	}
}
