package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Starting the consumer of rabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		fmt.Printf("Error in connecting the rabbitMQ : %v", err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Successfully established the connection to the rabbitMQ server")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("Error in channel of the rabbitMQ : %v", err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("Error in consuming the messages from rabbitMQ : %v", err)
		panic(err)
	}

	blockCh := make(chan bool) // This channel will block the main function from exiting before all megs recieved
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message : %s - %v = %s\n ", d.Body, d.Timestamp, d.MessageId)
		}
	}()

	fmt.Println(" - waiting for messages - ")
	<-blockCh
}
