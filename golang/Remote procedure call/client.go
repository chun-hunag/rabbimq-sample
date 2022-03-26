package main

import (
	"./helper"
	"./rabbitMQService"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	queueName := mqService.TempQueueDeclare()
	deliveryChan := mqService.Consume(queueName, "rpc_consumer")

	corrId := randomString(32)
	n := bodyFrom(os.Args)

	builder := rabbitMQService.PublishBuilder{}
	builder.Init().
		SetContentType("text/plain").
		SetReplyTo(queueName).
		SetCorrelationId(corrId).
		SetBody([]byte(strconv.Itoa(n)))

	mqService.PublishByPublishing("", "rpc_queue", builder.Build())

	for d := range deliveryChan {
		if corrId == d.CorrelationId {
			res, err := strconv.Atoi(string(d.Body))
			helper.FailOnError(err, "Failed to convert body to integer")
			fmt.Println(fmt.Sprintf("%d 's answer: %d", n, res))
			break
		}
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	helper.FailOnError(err, "Failed to convert arg to integer")
	return n
}
