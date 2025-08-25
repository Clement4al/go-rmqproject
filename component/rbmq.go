package component

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// this function to handle errors
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func StartRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//  Declare test-queue
	q1, err := ch.QueueDeclare(
		"test-queue", // queue name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare test-queue")

	// Declare queue-pe
	q2, err := ch.QueueDeclare(
		"queue-pe-monioring", // queue name
		true,                 // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	failOnError(err, "Failed to declare pe-monitoring queue")
	//listers 1 n 2›
	msgs1, err := ch.Consume(
		q1.Name, // queue
		"",      // consumer tag
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register a consumer for test-queue")

	msgs2, err := ch.Consume(
		q2.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer for pe-monitoring")

	// Process messages from the queues
	send := make(chan bool)

	go func() {
		for d := range msgs1 {
			fmt.Printf(" [test-queue] Received: %s\n", d.Body)
		}
	}()

	go func() {
		for d := range msgs2 {
			fmt.Printf(" [queue-pe-monitoring] Received: %s\n", d.Body)
		}
	}()

	fmt.Println(" [*] Waiting for messages from both queues. To exit press CTRL+C")
	<-send
}
