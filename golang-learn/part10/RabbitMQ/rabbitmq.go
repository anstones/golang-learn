package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)


const MQURL = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string
	Exchange string
	key string
	Mqurl string
}

//创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string)*RabbitMQ  {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, key: key, Mqurl: MQURL}
}

// 断开channel和connection

func (r *RabbitMQ) Destory()  {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理
func (r *RabbitMQ)failOnErr(err error, msg string)  {
	if err != nil{
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func NewRabbitMQSimple(queueName string) *RabbitMQ  {
	// 创建rmq实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabb"+ "itmq!")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}


//队列生产消息
func (r *RabbitMQ)PublishSimple(msg string)  {
	// 1. 申请队列， 如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
		)
	if err != nil{
		fmt.Println(err)
	}

	// 调用channel 发送消息到队列

	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(msg),
			},
		)
}

// simple模式下消费者
func (r *RabbitMQ)ConsumeSimple()  {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
		)
	if err != nil{
		fmt.Println(err)
	}
	msgs ,err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
		)
	if err != nil{
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs{
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<- forever
}

