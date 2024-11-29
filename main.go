package main

import (
    "log"
    "go_timezone/config"
    "go_timezone/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    db, err := config.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }
    defer db.Close()

    router := gin.Default()

    router.GET("/current-time", handlers.GetTorontoTime(db))

    log.Println("Starting server on port 8080...")
    router.Run(":8080")
}
