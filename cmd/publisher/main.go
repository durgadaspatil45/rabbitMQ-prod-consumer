// This service publish the messages to the rabbitMQ
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("welcome to rabbitmq producer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("Error in connection to rmq %v", err)
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Successfully connected to rabbitMQ")

	ch, err := conn.Channel() //concurrent server channels to process the bulk of messages
	if err != nil {
		fmt.Printf("Error in initializing channel in rmq %v", err)
		panic(err)
	}
	defer ch.Close()

	/* The queue name may be empty, in which case the server will generate a unique name which will be returned in the Name field of Queue struct.
	Durable and Non-Auto-Deleted queues will survive server restarts and remain when there are no remaining consumers or bindings. Persistent publishings will be
	restored in this queue on server restart.These queues are only able to be bound to durable exchanges.Non-Durable and Auto-Deleted queues will not be redeclared
	on server restart and will be deleted by the server after a short time when the last consumer is canceled or the last consumer's channel is closed. Queues with this
	lifetime can also be deleted normally with QueueDelete. These durable queues can only be bound to non-durable exchanges.Non-Durable and Non-Auto-Deleted queues will
	remain declared as long as the server is running regardless of how many consumers. This lifetime is useful for temporary topologies that may have long delays between
	consumer activity. These queues can only be bound to non-durable exchanges.Durable and Auto-Deleted queues will be restored on server restart, but without active consumers
	will not survive and be removed. This Lifetime is unlikely to be useful.Exclusive queues are only accessible by the connection that declares them and will be deleted when
	the connection closes. Channels on other connections will receive an error when attempting to declare, bind, consume, purge or delete a queue with the same name.
	When noWait is true, the queue will assume to be declared on the server. A channel exception will arrive if the conditions are met for existing queues or attempting to modify
	an existing queue from a different connection.When the error return value is not nil, you can assume the queue could not be declared with these parameters, and the channel will be closed.*/

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	) //QueueDeclare declare the queue to hold the messages (we have to specify the name of the queue) and deliver to consumer
	if err != nil {
		fmt.Printf("Error in declaring queue in rmq %v", err)
		panic(err)
	}

	fmt.Println("Printing the queue", q)

	/* Since publishings are asynchronous, any undeliverable message will get returned by the server. Add a listener with Channel.NotifyReturn to handle any undeliverable message
	when calling publish with either the mandatory or immediate parameters as true. It is possible for publishing to not reach the broker if the underlying socket is shut down without
	pending publishing packets being flushed from the kernel buffers. The easy way of making it probable that all publishings reach the server is to always call Connection.Close
	before terminating your publishing application. The way to ensure that all publishings reach the server is to add a listener to Channel.NotifyPublish and put the channel in confirm
	mode with Channel.Confirm. Publishing delivery tags and their corresponding confirmations start at 1. Exit when all publishings are confirmed.*/

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			err = ch.Publish(
				"",
				"TestQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte("Publishing the messages to Rabbit queue"),
					Timestamp:   time.Now(),
					MessageId:   strconv.Itoa(i),
				},
			) // Important: The Key in Publish should have to be the Queue name
			if err != nil {
				fmt.Printf("Error in publishing message to queue in rmq %v", err)
				panic(err)
			}
			wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println("Successfully published the message to queue")
}
