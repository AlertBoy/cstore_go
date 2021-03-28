package middleware

import (
	"github.com/streadway/amqp"
	"log"
)

func Consume() {

	conn, err := amqp.Dial(RMQADDR)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	forever := make(chan bool)
	for routine := 0; routine < CONSUMERCNT; routine++ {
		go func(routineNum int) {
			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			q, err := ch.QueueDeclare(
				QUEUENAME,
				true, //durable
				false,
				false,
				false,
				nil,
			)

			failOnError(err, "Failed to declare a queue")

			msgs, err := ch.Consume(
				q.Name,
				"MsgWorkConsumer",
				false, //Auto Ack
				false,
				false,
				false,
				nil,
			)

			if err != nil {
				log.Fatal(err)
			}

			for msg := range msgs {
				log.Printf("In %d consume a message: %s\n", 0, msg.Body)
				log.Printf("Done")
				msg.Ack(false) //Ack
			}

		}(routine)
	}
	<-forever
}
