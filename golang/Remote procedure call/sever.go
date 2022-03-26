package main

import (
	"./helper"
	"./rabbitMQService"
	"log"
	"strconv"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	mqService.QueueDeclare("rpc_queue", false)
	mqService.SetQos(1, 0)
	deliveryChan := mqService.Consume("rpc_queue", "")

	forever := make(chan bool)

	go func() {
		for d := range deliveryChan {
			n, err := strconv.Atoi(string(d.Body))
			helper.FailOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)
			publishing := rabbitMQService.PublishBuilder{}
			publishing.Init().
				SetContentType("test/plain").
				SetCorrelationId(d.CorrelationId).
				SetBody([]byte(strconv.Itoa(response)))
			mqService.PublishByPublishing("", d.ReplyTo, publishing.Build())
			err = d.Ack(false)
			helper.FailOnError(err, "Failed to ack a message")
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
