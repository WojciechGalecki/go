package main

import (
	"example.com/app/db"
	"example.com/app/routes"
	"github.com/gin-gonic/gin"
)

const (
	serverPort = ":8080"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", routes.GetEvents)
	server.GET("/events/:id", routes.GetEvent)
	server.POST("/events", routes.CreateEvent)
	server.PUT("/events/:id", routes.UpdateEvent)
	server.DELETE("/events/:id", routes.DeleteEvent)

  server.POST("/signup", routes.Signup)
  server.POST("/login", routes.Login)

	server.Run(serverPort)
}
