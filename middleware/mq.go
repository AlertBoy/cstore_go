package middleware

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

const (
	RMQADDR     = "amqp://root:123456@localhost:5672/"
	QUEUENAME   = "queForTest"
	PRODUCERCNT = 5
	CONSUMERCNT = 20
)

func Generate() {

	conn, err := amqp.Dial(RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(PRODUCERCNT)

	for routine := 0; routine < PRODUCERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				QUEUENAME, //Queue name
				true,      //durable
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "Failed to declare a queue")

			for i := 0; i < 500; i++ {
				msgBody := fmt.Sprintf("Message_%d_%d", routineNum, i)

				err = ch.Publish(
					"",     //exchange
					q.Name, //routing key
					false,
					false,
					amqp.Publishing{
						DeliveryMode: amqp.Persistent, //Msg set as persistent
						ContentType:  "text/plain",
						Body:         []byte(msgBody),
					})

				log.Printf(" [x] Sent %s", msgBody)
				failOnError(err, "Failed to publish a message")
			}

			wg.Done()
		}(routine)
	}

	wg.Wait()

	log.Println("All messages sent!!!!")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
