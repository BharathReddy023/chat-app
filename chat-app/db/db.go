// db/db.go

package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    
    connStr := "postgres://bharath:Password@123@localhost/chat?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return db, nil
}
