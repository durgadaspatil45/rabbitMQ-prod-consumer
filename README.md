# rabbitMQ-Publisher-Consumer

## Publisher

  - These are following steps need to follow to publish the messages to the rabbitMQ server
    - run the rabbitMQ container with following command
    - establis the connection to the RMQ server (amqp.Dial("amqp://username:password@fqdn:5672"))
    - default username - guest and password - guest
    - create the concurrent server channel to process the bulk of messages - conn.Channel()
    - create the Queue  to store the messages so that consumer can consume the messages from the channel - ch.QueueDeclare()
    - publish the messages to the Queue - ch.Publish()


## Consumer

  - These are following steps need to follow to consume the messages from the rabbitMQ server
    - establis the connection to the RMQ server (amqp.Dial("amqp://username:password@fqdn:5672"))
    - default username - guest and password - guest
    - create the concurrent server channel to process the bulk of messages - conn.Channel()
    - consume the messages from the Queue - ch.Publish()

### **To run the service**

  - docker run -d --hostname my-rabbit --name rabbit-server -p 15672:15672 -p 5672:5672 rabbitmq:3-management

  - go run cmd/publisher/main.go

  - go run cmd/consumer/main.go
