package util

import (
	"log"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var RabbitConnection *amqp.Connection
var RabbitChannel *amqp.Channel

// 初始化RabbitMQ
func InitRabbitMQ() {
	host := viper.GetString("rabbitmq.host")
	port := viper.GetString("rabbitmq.port")
	username := viper.GetString("rabbitmq.username")
	password := viper.GetString("rabbitmq.password")
	url := "amqp://" + username + ":" + password + "@" + host + ":" + port + "/"
	conn, err := amqp.Dial(url)
	failOnError(err, "[RabbitMQ]服务连接异常")
	RabbitConnection = conn
	ch, err := conn.Channel()
	failOnError(err, "[RabbitMQ]通道连接异常")
	RabbitChannel = ch
}

// 发送信息
func SendMessage(ch *amqp.Channel, message string) {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "[RabbitMQ]队列连接异常")
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "[RabbitMQ]消息传输异常")
	log.Printf("[RabbitMQ]发送消息：%s\n", message)
}

// 接收信息
func GetMessage(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "[RabbitMQ]连接异常")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "[RabbitMQ]消费注册异常")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("[RabbitMQ]等待消息。推出请按CTRL+C。")
	<-forever
}

// 错误处理
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// 关闭RabbitMQ的connection和channel
func CloseRabbitMQ() {
	RabbitConnection.Close()
	RabbitChannel.Close()
}
