package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"products_ms/src/consumers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/consume", ConsumeHandler).Methods("GET")

	http.Handle("/", r)

	// Start a Goroutine to consume messages from RabbitMQ
	go ConsumeMessages()

	http.ListenAndServe(":3000", nil)
}

func ConsumeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle API request here
	fmt.Fprintln(w, "API Endpoint: Consume RabbitMQ Messages")
}

func ConsumeMessages() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://root:1234@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()
	
	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	// defer ch.Close()
	
	consumers.Consumers(ch)
	
}
