package actions

import (

	"github.com/gin-gonic/gin"

	database "github.com/datasektionen/facebook/server/db"
)


func CreateSchedule(c *gin.Context) {

	database.InitDB()
    db := database.GetDB()

    db.Create(&database.SCHEDULE{
        Key: "test", 
        Overseers: "Nils Malmberg", 
        Comments: "Detta är en väldigt seriös kommentar som bör läsas",
		Checklist: 1,
    })

}