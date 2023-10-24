package send

import (
	"context"
	"fmt"
	"log"

	//   "time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// func main() {
// 	conn, err := ConnectToRabbit()
// 	failOnError(err, "Failed connection")

// 	ch, err := CreateChRabbit(conn)
// 	failOnError(err, "Failed to open a channel")

// 	q, err := CreateQueue(ch)
// 	failOnError(err, "Failed to declare a queue")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	body := "Hello World!"
// 	err = SendToQueue([]byte(body), ctx, ch, q)
// 	failOnError(err, "Failed to publish a message")
// 	log.Printf(" [x] Sent %s\n", body)
// }

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}

func ConnectToRabbit(host string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:5672/", host))
	if err != nil  {
		return nil, err
	}
	return conn, nil
}

func CreateChRabbit(c *amqp.Connection) (*amqp.Channel, error) {
	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func SendToQueue(body []byte, ctx context.Context, ch *amqp.Channel) error {
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		"hello", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		ContentType: "text/plain",
		Body:        body,
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", body)
	return nil
}

func CreateQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
}
