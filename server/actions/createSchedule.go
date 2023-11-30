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

func CreateSchedule(c *gin.Context) {

	database.InitDB()
    db := database.GetDB()
    
    // get data from checklist
    var checklistItems []database.CHECKLIST
	if err := db.Preload(clause.Associations).Find(&checklistItems).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
    var parsedChecklistItems []database.ChecklistItem
    for _, item := range checklistItems {
        var checklistItem database.ChecklistItem
        if err := json.Unmarshal([]byte(item.ChecklistItem), &checklistItem); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

        parsedChecklistItems = append(parsedChecklistItems, checklistItem)
    }

    var extractedChecklistIDs = make(map[string]bool)

    // Iterate over parsedChecklistItems to extract ChecklistItemID
    for _, item := range parsedChecklistItems {
        for _, swedishValue := range item.SwedishValues {
            extractedChecklistIDs[swedishValue.ChecklistItemID] = false
        }
    
        for _, englishValue := range item.EnglishValues {
            extractedChecklistIDs[englishValue.ChecklistItemID] = false
        }
    }

    fmt.Println(extractedChecklistIDs)

    // Convert the map to a JSON string
    jsonString, err := json.Marshal(extractedChecklistIDs)
    if err != nil {
        // Handle the error, for example, log it or return an error response
        fmt.Println("Error marshaling to JSON:", err)
        return
    }

    // Now jsonString contains the JSON representation of extractedChecklistIDs as a string
    fmt.Println(string(jsonString))

    data_entry := &database.SCHEDULE{
        Key: "thisisakey", 
        Overseers: "name name2", 
        Comments: "Detta är en väldigt seriös kommentar som bör läsas",
        ChecklistJSON: string(jsonString),
    }

    err = db.Create(&data_entry).Error
    if err != nil {
        c.JSON(500, gin.H{
            "error": err,
        })
        return
    }

    c.JSON(200, gin.H{
		"message": "Schedule items inserted successfully",
	})
}