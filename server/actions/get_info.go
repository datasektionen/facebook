package actions

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	database "github.com/datasektionen/facebook/server/db"
)


func GetInfo(c *gin.Context) {
    database.InitDB()
    db := database.GetDB()

    // Define a variable to store the retrieved data
    var schedule []database.SCHEDULE

    // Query the "SCHEDULE" table and retrieve all records
    if err := db.Find(&schedule).Error; err != nil {
        // Handle the error if there's any
        fmt.Println("Error querying database:", err)
        c.String(http.StatusInternalServerError, "Internal Server Error")
        return
    }

    // Print out the retrieved data
    for _, schedule := range schedule {
        fmt.Printf("Nyckel: %s\n", schedule.Key)
    }

    fmt.Println("Hello world!")

    c.String(http.StatusOK, "")
}