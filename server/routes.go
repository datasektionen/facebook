package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"

	"github.com/datasektionen/facebook/server/actions"
	"github.com/datasektionen/facebook/server/db"
)

func InitRoutes(r *gin.RouterGroup) {
	db.InitDB()

	// r.Use(cors.New(cors.Options{}))

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	r.GET("/trigger", actions.CreateChecklist)
	r.GET("/send_websocket", actions.SendWebSocketMessageToClients)
}

func InitWebsocket(r *gin.RouterGroup){
	r.GET("", actions.SendWebsocket)
}

