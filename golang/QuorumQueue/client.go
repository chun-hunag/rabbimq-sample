package main

import (
	"./helper"
	"./rabbitMQService"
	"os"
	"strconv"
)

func main() {
	n := getArgs(os.Args)
	mqService := rabbitMQService.NewRabbitMQService()

	for i := 0; i < n; i++ {
		mqService.Publish("quorum.exchange", "quorum.queue", strconv.Itoa(i))
	}
}

func getArgs(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		helper.FailOnError(nil, "wrong or absent Args")
	} else {
		s = args[1]
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		helper.FailOnError(err, "Args not integer")
	}

	return n
}
