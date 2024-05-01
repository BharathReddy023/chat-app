package handlers

import (
	"chat/db" // Import your database package
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert user into the database
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		user.Username, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Fetch user from database
	row := db.QueryRow("SELECT id, password FROM users WHERE email = $1", user.Email)
	var id int
	var hashedPassword string
	err = row.Scan(&id, &hashedPassword)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})
	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return JWT token in response
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query all users from the database
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Iterate over the rows and collect user data
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Return user data in the response
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from request parameters
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert user ID to integer
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := db.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute SQL DELETE statement
	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User with ID %d deleted successfully", id)
}
