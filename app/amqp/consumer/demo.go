package consumer

import (
	"log"
	"my-app/library/middleware/rabbitmq"
)

func init()  {
	consumer()
}

func consumer() {
	rabbitmq.Debug = true

	connStr := rabbitmq.MidAmpq.MqConnStr("default")

	conn, err := rabbitmq.Dial(connStr)
	if err != nil {
		log.Panic(err)
	}

	queueName := "test-queue"

	consumeCh, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	go func() {
		d, err := consumeCh.Consume(queueName, "", false, false, false, false, nil)
		if err != nil {
			log.Panic(err)
		}

		for msg := range d {
			log.Printf("msg: %s", string(msg.Body))
			msg.Ack(true)
		}
	}()

}
