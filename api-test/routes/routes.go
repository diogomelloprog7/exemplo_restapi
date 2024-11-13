package routes

import (
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events/:id")
	server.GET("/events", getEvents)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events", middleware.Authenticate, createEvent)
	server.POST("/signup", signup)
	server.POST("/events/:id", updateEvent)
	server.POST("/login", login)

}
