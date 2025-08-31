package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		// Public routes (no authentication required)
		v1.POST("/auth/register", app.RegisterUser)
		v1.POST("/auth/login", app.LoginUser)

		// Protected routes (authentication required)
		protected := v1.Group("")
		protected.Use(AuthMiddleware(app.jwtSecret))
		{
			protected.POST("/users", app.CreateUser)
			protected.GET("/users", app.GetUsers)
			protected.GET("/users/:id", app.GetUser)
			protected.PUT("/users/:id", app.UpdateUser)
			protected.DELETE("/users/:id", app.DeleteUser)

			protected.POST("/events", app.CreateEvent)
			protected.GET("/events", app.GetEvents)
			protected.GET("/events/:id", app.GetEvent)
			protected.PUT("/events/:id", app.UpdateEvent)
			protected.DELETE("/events/:id", app.DeleteEvent)
			protected.POST("/events/:id/attendees/:user_id", app.AddAttendeeToEvent)
			protected.GET("/events/:id/attendees", app.GetAttendeesForEvent)

			protected.POST("/attendees", app.CreateAttendee)
			protected.GET("/attendees", app.GetAttendees)
			protected.GET("/attendees/:id", app.GetAttendee)
			protected.PUT("/attendees/:id", app.UpdateAttendee)
			protected.DELETE("/attendees/:id", app.DeleteAttendee)
		}
	}

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return g
}
