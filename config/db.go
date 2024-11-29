package config

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    host := "localhost"
    port := 5432
    user := "postgres"
    password := "Prince@107040"
    dbname := "toronto_time"

    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    fmt.Println("Database connection successful")
    return db, nil
}
