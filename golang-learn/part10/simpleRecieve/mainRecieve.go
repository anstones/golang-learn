package main

import "golang-learn/part10/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(""+"kuteng")
	rabbitmq.ConsumeSimple()
}