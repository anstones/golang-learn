package main

import (
	"fmt"
	"golang-learn/part10/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(""+"kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功")
}


