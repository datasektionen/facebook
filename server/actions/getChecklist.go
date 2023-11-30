package actions

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
    "fmt"
    "encoding/json"


	database "github.com/datasektionen/facebook/server/db"
)

func GetChecklist(c *gin.Context) {
	db := database.GetDB()

    fmt.Println("Gathering data")

	var checklistItems []database.CHECKLIST
	if err := db.Preload(clause.Associations).Find(&checklistItems).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

    // Now, loop through each item and parse the JSON
    var parsedChecklistItems []database.ChecklistItem
    for _, item := range checklistItems {
        var checklistItem database.ChecklistItem

        // Assuming ChecklistItem field in CHECKLIST contains the JSON data
        if err := json.Unmarshal([]byte(item.ChecklistItem), &checklistItem); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

        parsedChecklistItems = append(parsedChecklistItems, checklistItem)
    }

    c.JSON(200, parsedChecklistItems)

    fmt.Println("Data gathered")

}