package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	database "github.com/datasektionen/facebook/server/db"

	server "github.com/datasektionen/facebook/server"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	database.InitDB()

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.Static("/public", "./public")

	// Add CORS middleware
	r.Use(corsMiddleware())

	api := r.Group("/api")
	server.InitRoutes(api)

	websocket := r.Group("/websocket")
	server.InitWebsocket(websocket)

	r.Run(fmt.Sprintf(":%d", 5001))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
