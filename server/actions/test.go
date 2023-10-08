package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// database "github.com/datasektionen/facebook/server/db"
)

func Test(c *gin.Context) {
	// db := database.GetDB()

	c.String(http.StatusOK, "")
}