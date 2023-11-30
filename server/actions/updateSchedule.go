package actions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
    "encoding/json"
    "fmt"

	database "github.com/datasektionen/facebook/server/db"
)

type ChecklistJSON struct {
    Key string
    Value bool
}

func UpdateSchedule(c *gin.Context) {

	database.InitDB()
    db := database.GetDB()
    
    // get data from checklist
    var scheduleItems []database.SCHEDULE
	if err := db.Preload(clause.Associations).Find(&scheduleItems).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}  

	


    c.JSON(200, gin.H{
		"message": "Schedule items inserted successfully",
	})
}