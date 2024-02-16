package actions

import (
	"fmt"

	"github.com/gin-gonic/gin"
	// "encoding/json"
)

func UpdateSchedule(c *gin.Context) {

	data := c.Query("key")

	fmt.Println("THIS IS COOL DATA: ", data)

	SendWebSocketDataToClients(c)

	c.JSON(200, gin.H{
		"data": data,
	})
}
