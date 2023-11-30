package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// cors "github.com/rs/cors/wrapper/gin"

	actions "github.com/datasektionen/facebook/server/actions"
)

func InitRoutes(r *gin.RouterGroup) {
	// r.Use(cors.New(cors.Options{}))

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	r.GET("/trigger", actions.CreateChecklist)
	r.GET("/send_websocket", actions.SendWebSocketMessageToClients)
	// r.GET("/create_checklist", actions.CreateChecklist)  // This is only for updating/ recreating the entire checklist. This should not be open by default for obvoius reasons 
	r.GET("/create_schedule", actions.CreateSchedule)

	r.PUT("/update_schedule", actions.UpdateSchedule)
	r.GET("/test", actions.SendWebSocketDataToClients)
}

func InitWebsocket(r *gin.RouterGroup){
	r.GET("", actions.SendWebsocket)
}

