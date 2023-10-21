package consumers

import (
	"encoding/json"
	"fmt"
	"log"

	"products_ms/src/config"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type Product struct {
	Title string
	Price string
}

type PayloadMessage struct {
	Action string
	Data string
}

type MessageConsumed struct {
	Payload PayloadMessage
}

func productCb(messages <-chan amqp.Delivery) {
	for msg := range messages {
		receivedMessage := MessageConsumed{}
		err := json.Unmarshal(msg.Body, &receivedMessage)
		if err != nil {
			log.Println(err)
		}
		switch receivedMessage.Payload.Action {
		case "new":
			newProduct([]byte(receivedMessage.Payload.Data))
		case "otherType":
			// Handle other types similarly
		default:
			fmt.Println("Unknown Action")
		}
		
	}
}

func newProduct(payloadData []byte) {
	db, errDbCon := database.ConnectDB()
	newUUID := uuid.New()

	if errDbCon != nil {
		log.Fatal("DB Not Connected: Product.New")
	}
	product := Product{}
	err := json.Unmarshal(payloadData, &product)
	if err != nil {
		log.Println(err)
	}
	sqlStatement := `INSERT INTO product (id) VALUES ($1)`
	_, err = db.Exec(sqlStatement, newUUID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("product => ",product)
	defer db.Close()
}