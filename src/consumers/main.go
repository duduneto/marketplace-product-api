package consumers

import (
	"log"

	"github.com/streadway/amqp"
)

func Consumers(channel *amqp.Channel) {
    queues := []string{"products", "products.stock"}
	
	for _, queueName := range queues {
		go consumerHandler(channel, queueName, productCb)
	}
}

func getDeliveriesChannel(channel *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		queueName,
		"",
		true,  // auto-acknowledge messages
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
}

func consumerHandler(
	channel *amqp.Channel,
	queueName string,
	callback func(<-chan amqp.Delivery),
) {
	msgs, err := channel.Consume(
		queueName,
		"",
		true,  // auto-acknowledge messages
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	// Start consuming messages
	productCb(msgs)
}
