package main

import (
	"./rabbitMQService"
)

func main() {
	mqService := rabbitMQService.NewRabbitMQService()
	mqService.ExchangeDeclare("logs", rabbitMQService.Fanout)
	queueName := mqService.TempQueueDeclare()
	mqService.QueueBind(queueName, "", "logs")
	mqService.Consume(queueName, "")
}
