// main.go

package main

import (
	"chat/handlers"
	"chat/kafka"
	"chat/redis"
	"net/http"
	"time"
)

func main() {

	type Message struct {
		ID         int       `json:"id"`
		SenderID   int       `json:"sender_id"`
		ReceiverID int       `json:"receiver_id"`
		Text       string    `json:"text"`
		SentAt     time.Time `json:"sent_at"`
	}

	// Initialize Kafka producer
	kafka.InitProducer()
	defer kafka.Producer.Close()

	// Initialize Redis client
	redis.InitClient()

	http.HandleFunc("/api/auth/register", handlers.Register)
	http.HandleFunc("/api/auth/login", handlers.Login)
	http.HandleFunc("/api/users/get", handlers.GetUsers)
	http.HandleFunc("/api/users/delete", handlers.DeleteUser)
	http.HandleFunc("/api/messages/send", handlers.SendMessage)
	http.HandleFunc("/api/messages/history", handlers.GetMessageHistory)

	http.ListenAndServe(":8080", nil)
}
