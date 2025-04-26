package routes

import (
	"example.com/udemy_course/middlewares"
	"github.com/gin-gonic/gin"
)


func InitializeRoute(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	
	server.POST("/sign-up", signUp)
	server.POST("/sign-in", signIn)

	server.POST("/events/:id/register", middlewares.Authenticate, register)
	server.DELETE("/events/:id/cancel", middlewares.Authenticate, cancel)
}
