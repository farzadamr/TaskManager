package api

import (
	"github.com/farzadamr/TaskManager/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, handlrs ...interface{}) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	var (
		authHandler *handlers.AuthHandler
		taskHandler *handlers.TaskHandler
	)

	for _, handler := range handlrs {
		switch h := handler.(type) {
		case *handlers.AuthHandler:
			authHandler = h
		case *handlers.TaskHandler:
			taskHandler = h
		}
	}

	e.POST("api/v1/register", authHandler.Register)
	e.POST("api/v1/login", authHandler.Login)

	api := e.Group("/api/v1")

	taskRoutes := api.Group("/tasks")
	taskRoutes.POST("", taskHandler.CreateTask)
	taskRoutes.GET("", taskHandler.GetAllTasks)
	taskRoutes.GET("/:id", taskHandler.GetTask)
	taskRoutes.PUT("/:id", taskHandler.UpdateTask)
	taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	taskRoutes.PATCH("/:id/complete", taskHandler.CompleteTask)
}
