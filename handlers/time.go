package handlers

import (
    "database/sql"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

func GetTorontoTime(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get current time in Toronto timezone
        location, err := time.LoadLocation("America/Toronto")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load timezone"})
            return
        }

        currentTime := time.Now().In(location)

        // Log the time in the database
        _, err = db.Exec("INSERT INTO time_log (timestamp) VALUES ($1)", currentTime)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log time in database"})
            return
        }

        // Return time in JSON format
        c.JSON(http.StatusOK, gin.H{
            "current_time": currentTime.Format(time.RFC3339),
        })
    }
}
