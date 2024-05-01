package handlers

import (
	"chat/db"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Text       string    `json:"text"`
	SentAt     time.Time `json:"sent_at"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate message fields (e.g., check if sender and receiver IDs are valid)

	// Insert message into the database
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO messages (sender_id, receiver_id, text) VALUES ($1, $2, $3)",
		message.SenderID, message.ReceiverID, message.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	// Parse sender and receiver IDs from the request query parameters
	senderIDStr := r.URL.Query().Get("sender_id")
	receiverIDStr := r.URL.Query().Get("receiver_id")

	senderID, err := strconv.Atoi(senderIDStr)
	if err != nil {
		http.Error(w, "Invalid sender ID", http.StatusBadRequest)
		return
	}

	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		http.Error(w, "Invalid receiver ID", http.StatusBadRequest)
		return
	}

	// Fetch message history from the database
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT sender_id, receiver_id, text, sent_at FROM messages WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1) ORDER BY sent_at ASC",
		senderID, receiverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.SenderID, &message.ReceiverID, &message.Text, &message.SentAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}

	// Return message history in the response
	json.NewEncoder(w).Encode(messages)
}
