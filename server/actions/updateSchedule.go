package actions

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	// "encoding/json"

)

func UpdateSchedule(c *gin.Context) {

	data := GetData(c.Query("key"))




    c.JSON(200,gin.H{
        "data": data,
    })
}