package database

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    fmt.Println("Connected to the database")
    return db, nil
}
