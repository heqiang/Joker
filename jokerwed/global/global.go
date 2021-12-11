package global

import "github.com/streadway/amqp"

var (
	MqConn *amqp.Connection
	MqChan *amqp.Channel
)
