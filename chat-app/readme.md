
# Real-Time Chat Application

 Welcome to the Real-Time Chat Application! This application allows users to engage in real-time private messaging with each other. It is built with scalability and efficiency in mind, utilizing technologies such as Kafka, Redis, and GoLang.


## Introduction

This application allows users to:

Create user accounts
Send and receive private messages in real-time
Retrieve message history
Manage user sessions using Redis
Utilize Kafka for real-time message delivery

## Technologies Used
Database: Postgres
Caching: Redis
Messaging Queue: Kafka
Backend Development: GoLang
Testing:  Go testing tools (GoLang)
Performance Testing: Jmeter

## Getting Started

Clone the Repository:
git clone git@github.com:BharathReddy023/chat-app.git
cd yourrepository

Install Dependencies:
Ensure you have Postgres, Redis, Kafka, GoLang installed on your machine.
GoLang dependencies using go mod tidy.

Database Setup:
Create a Postgres database and configure the connection in your application.
Create the required tables using the provided SQL queries.

### queries:

to create database:
create database chat;

to connect to database:
\c chat

to create tables use below commands:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL,
    receiver_id INT NOT NULL,
    text TEXT NOT NULL,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (receiver_id) REFERENCES users(id)
);
Configuration:
Set up environment variables for configuration parameters such as database connection details, Kafka configuration, etc.
Run the Application:
go run main.go (for GoLang).

Access the Application:
The application will be accessible at http://localhost:8080.

## Endpoints
The Real-Time Chat Application provides the following endpoints:

User Management:
/api/auth/register: Register a new user.
/api/auth/login: Log in an existing user.
/api/users/get: Get users information.
/api/users/delete: Delete a user account.
Messaging:
/api/messages/send: Send a private message.
/api/messages/history: Retrieve message history between two users.

to test with postman use below syntax for body with method post:

## register
{
    "username": "user1",
    "email": "user1@example.com",
    "password": "password123"
}


## login
{
    "email": "user1@example.com",
    "password": "password123"
}

## messages
{
    "sender_id": 2,
    "receiver_id": 1,
    "text": "good , how are you?"
}

## Testing:
create a new file named as $BASEFILE_test.go
 write testcases

  To run unit tests for standalone functions, use the following command:
 go test

  to check the code coverage
 go test -cover

 ### Unit Tests
Unit tests are located in the *_test.go files adjacent to the source files they test.
### Integration Tests with PostgreSQL
Integration tests interact directly with the PostgreSQL database to confirm that CRUD operations perform as intended. Before running integration tests, ensure you have PostgreSQL installed and running on your machine. Update the database connection string in the test files accordingly.
 ### API Endpoint Tests
API endpoint tests validate the correct behavior of each endpoint for both valid and invalid request scenarios. You can use the Go testing framework or a tool like Postman to execute these tests.

## Postman Collection

### Export Collection
1. Open Postman.
2. Click on "Collections" in the sidebar.
3. Hover over the collection you want to export and click on the ellipsis (three dots).
4. Select "Export" and choose the format "Collection v2.1".
5. Save the exported collection file to your desired location.

### Import Collection
1. Open Postman.
2. Click on "Import" in the top left corner.
3. Choose the exported collection file.
4. Click on "Import" to add the collection to your Postman workspace.

Now you can use the exported Postman collection to test your API endpoints easily.


### Learnings:
- **Go Fundamentals:** Mastered basics like syntax, data types, and control structures.
- **REST Principles:** Understood URL design, HTTP methods, and status codes.
- **CRUD Operations:** Implemented Create, Read, Update, Delete operations efficiently.
- **Error Handling:** Learned to manage errors effectively for a robust API.
- **Documentation:** Created comprehensive documentation for clear usage.
- **Testing:** Provided Postman collection for easy API testing.