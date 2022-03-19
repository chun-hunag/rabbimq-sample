package main

import (
	"./rabbitMQService"
	"os"
	"strings"
)

func main() {
	body := bodyFrom(os.Args)
	mqService := rabbitMQService.NewRabbitMQService()
	mqService.QueueDeclare("queue.test")
	mqService.Publish("queue.test", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
