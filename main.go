package main

import (
	"example.com/app/db"
	"example.com/app/middlewares"
	"example.com/app/routes"
	"github.com/gin-gonic/gin"
)

const serverPort = ":8080"

func main() {
	db.InitDB()
	server := gin.Default()

    server.POST("/signup", routes.Signup)
	server.POST("/login", routes.Login)
	server.GET("/events", routes.GetEvents)
	server.GET("/events/:id", routes.GetEvent)

    authenticated := server.Group("/")
    authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", routes.CreateEvent)
	authenticated.PUT("/events/:id", routes.UpdateEvent)
	authenticated.DELETE("/events/:id", routes.DeleteEvent)
    authenticated.POST("/events/:id/register", routes.RegisterForEvent)
    authenticated.DELETE("/events/:id/register", routes.CancelRegistration)

	server.Run(serverPort)
}
