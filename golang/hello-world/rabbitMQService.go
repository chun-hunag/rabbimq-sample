package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQService struct {
	account    string
	password   string
	host       string
	vhost      string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQService() *RabbitMQService {
	// should be replaced by env config load
	account := "guest"
	password := "guest"
	host := "localhost"
	vhost := "/"
	rabbitMQService := &RabbitMQService{account: account, password: password, host: host, vhost: vhost}
	rabbitMQService.getConnection()
	rabbitMQService.getChannel()
	return rabbitMQService
}

func (r *RabbitMQService) getConnection() *amqp.Connection {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s%s", r.account, r.password, r.host, r.vhost))
	failOnError(err, "Failed to connect to RabbitMQ")
	r.connection = connection
	return r.connection
}

func (r *RabbitMQService) getChannel() *amqp.Channel {
	if r.connection == nil {
		log.Fatalf("%s", "Connection is nil.")
	}

	channel, err := r.connection.Channel()
	failOnError(err, "Failed to open a channel")
	r.channel = channel
	return r.channel
}

func (r *RabbitMQService) isReady() bool {
	return r.connection != nil && r.channel != nil
}

func (r *RabbitMQService) close() {
	if r.channel != nil {
		err := r.channel.Close()
		failOnError(err, "Failed to close channel.")
	}

	if r.connection != nil {
		err2 := r.connection.Close()
		failOnError(err2, "Failed to close connection.")
	}
}

func (r *RabbitMQService) QueueDeclare(name string) {
	if r.isReady() {
		_, err := r.channel.QueueDeclare(
			name,  // name
			false, // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")
	}
}

func (r *RabbitMQService) Publish(queueName, body string) {
	if r.isReady() {
		err := r.channel.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
		failOnError(err, "Failed to publish a message")
	}
}

func (r *RabbitMQService) Consume(queueName, consumerName string) {
	messages, err := r.channel.Consume(
		queueName,    // name
		consumerName, // consumer tags
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
