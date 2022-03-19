package rabbitMQService

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type ExchangeType string

const (
	Direct  ExchangeType = "direct"
	Topic   ExchangeType = "topic"
	Headers ExchangeType = "headers"
	Fanout  ExchangeType = "fanout"
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
	rabbitMQService.SetQos(1, 0)
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

func (r *RabbitMQService) ExchangeDeclare(name string, exchangeType ExchangeType) {
	err := r.channel.ExchangeDeclare(
		name,
		string(exchangeType),
		true,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to declare an exchange")
}

func (r *RabbitMQService) QueueDeclare(name string, durable bool) {
	if r.isReady() {
		_, err := r.channel.QueueDeclare(
			name,    // name
			durable, // durable
			false,   // delete when unused
			false,   // exclusive
			false,   // no-wait
			nil,     // arguments
		)
		failOnError(err, "Failed to declare a queue")
	}
}

func (r *RabbitMQService) TempQueueDeclare() string {
	if r.isReady() {
		queue, err := r.channel.QueueDeclare(
			"",
			false,
			false,
			true,
			false,
			nil,
		)
		failOnError(err, "Failed to declare a queue")
		return queue.Name
	}
	return ""
}

func (r *RabbitMQService) Publish(exchangeName, queueName, body string) {
	if r.isReady() {
		err := r.channel.Publish(
			exchangeName, // exchange
			queueName,    // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
			},
		)
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s", body)
	}
}

func (r *RabbitMQService) Consume(queueName, consumerName string) {
	messages, err := r.channel.Consume(
		queueName,    // name
		consumerName, // consumer tags
		false,        // auto-ack
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
			// simulate work time
			dotCount := bytes.Count(message.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			message.Ack(false) // manually acknowledgement
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func (r *RabbitMQService) QueueBind(queueName, routingKey, exchangeName string) {
	if r.isReady() {
		err := r.channel.QueueBind(
			queueName,
			routingKey,
			exchangeName,
			false,
			nil,
		)
		failOnError(err, "Failed to QueueBind")
	}
}

func (r *RabbitMQService) SetQos(count, size int) {
	err := r.channel.Qos(
		count,
		size,
		false,
	)
	failOnError(err, "Failed to set QoS")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
