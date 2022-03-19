package main

import (
	"./rabbitMQService"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	mqService.QueueDeclare("queue.test")
	mqService.Consume("queue.test", "consumer.test")
}
