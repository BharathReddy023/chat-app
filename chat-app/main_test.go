package main_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"chat/handlers"
)

var db *sql.DB

func TestRegisterUser(t *testing.T) {
	initDB()
	defer db.Close()

	// Create a new request to register a user
	userData := handlers.User{
		Username: "bharath",
		Email:    "bharath@gmail.com",
		Password: "password",
	}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function to register the user
	handlers.Register(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

func TestSendMessage(t *testing.T) {
	initDB()
	defer db.Close()

	// Insert sender and receiver users into the database for testing
	_, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", "sender", "sender@gmail.com", "password")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", "receiver", "receiver@gmail.com", "password")
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request to send a message
	messageData := handlers.Message{
		SenderID:   1,
		ReceiverID: 2,
		Text:       "Hello",
		SentAt:     time.Now(),
	}
	jsonData, err := json.Marshal(messageData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/messages/send", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function to send the message
	handlers.SendMessage(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	
}

// Initialize the database connection for testing
func initDB() {
	dbinfo := "user=bharath password=Password@123 dbname=chat sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}

