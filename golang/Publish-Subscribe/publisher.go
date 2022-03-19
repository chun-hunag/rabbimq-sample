package main

import (
	"./rabbitMQService"
	"os"
	"strings"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	mqService.ExchangeDeclare("logs", rabbitMQService.Fanout)
	mqService.Publish("logs", "", bodyFrom(os.Args))
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
