package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	server "github.com/datasektionen/facebook/server"
)

func main() {
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.Static("/public", "./public")

	api := r.Group("/api")
	server.InitRoutes(api)

	websocket := r.Group("/websocket")
	server.InitWebsocket(websocket)

	r.Run(fmt.Sprintf(":%d", 5001))
}