package consumers

import (
	"fmt"

	"github.com/streadway/amqp"
)

func productStockCb(messages <-chan amqp.Delivery) {
	for msg := range messages {
		message := string(msg.Body)
		fmt.Printf("Received a message: %s\n", message)
	}
}