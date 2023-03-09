package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// go 操作 rabbitMQ

// rabbitMQ的URL信息
// url:"guest://用户名：密码@ip:port/vhost"
const MQURL = "amqp://admin:admin@172.29.194.124:15672/my_vhost"

// 存放信息的结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	Mqurl string
}

// 通用方法
// 创建RabbiMQ实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	// 创建连接
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "连接错误")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq

}

// 断开channel和connection的连接释放资源
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 自定义错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
func main() {

}
