package main

func main() {
	mqService := NewRabbitMQService()
	mqService.QueueDeclare("queue.test")
	mqService.Publish("queue.test", "hello-world")
}
