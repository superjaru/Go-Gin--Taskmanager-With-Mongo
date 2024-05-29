package main

import (
	"github.com/gin-gonic/gin"

	"task-manager/handlers"
	"task-manager/middleware"
)

func main() {
	router := gin.Default()

	auth := router.Group("/api/auth")
	auth.POST("/register", handlers.Register)
	auth.POST("/login", handlers.Login)
	

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	api.GET("/tasks", handlers.ListTasks)
	api.POST("/tasks", handlers.CreateTask)
	api.GET("/tasks/:id", handlers.GetTask)
	api.PUT("/tasks/:id", handlers.UpdateTask)
	api.DELETE("/tasks/:id", handlers.DeleteTask)
	

	router.Run(":8080")
}
