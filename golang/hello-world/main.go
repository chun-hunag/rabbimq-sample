package main

func main() {
	mqService := NewRabbitMQService()
	//mqService.QueueDeclare("queue.test")
	//mqService.Publish("queue.test", "hello-world")

	mqService.Consume("queue.test", "consumer.test")
}
