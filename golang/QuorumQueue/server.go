package main

import (
	"./helper"
	"./rabbitMQService"
	"log"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	queueName := "quorum-queue"
	mqService.QuorumQueueDeclare(queueName)
	mqService.ExchangeDeclare("quorum.exchange", rabbitMQService.Fanout)
	mqService.QueueBind(queueName, "quorum.queue", "quorum.exchange")
	deliverChan := mqService.Consume(queueName, "client")
	forever := make(chan bool)
	go func() {
		for d := range deliverChan {
			message := string(d.Body)
			log.Printf("%s%s", "Consume message: ", message)
			err := d.Ack(false)
			helper.FailOnError(err, "Failed to ack a message")
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
