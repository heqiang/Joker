package initlize

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"spider/global"
	"spider/model"
	"spider/utils"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func InitMq() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			InsertMysql(d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InsertMysql(data []byte) {
	spiderItem := model.Article{
		ArticleId: utils.SnowFlake(),
	}
	err := json.Unmarshal(data, &spiderItem)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	tx := global.MySqlDb.Begin()
	res := tx.Create(&spiderItem)
	if res.RowsAffected < 1 {
		tx.Rollback()
		fmt.Println(res.Error.Error())
		return
	}
	tx.Commit()
}
