package api

import (
	"github.com/farzadamr/TaskManager/api/handlers"
	"github.com/farzadamr/TaskManager/services"
	"github.com/labstack/echo/v4"
 	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(e *echo.Echo, taskService services.TaskService) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")

	taskHandler := handlers.NewTaskHandler(taskService)

	taskRoutes := api.Group("/tasks")
	taskRoutes.POST("", taskHandler.CreateTask)
	taskRoutes.GET("", taskHandler.GetAllTasks)
	taskRoutes.GET("/:id", taskHandler.GetTask)
	taskRoutes.PUT("/:id", taskHandler.UpdateTask)
	taskRoutes.DELETE("/:id", taskHandler.DeleteTask)
	taskRoutes.PATCH("/:id/complete", taskHandler.CompleteTask)
}
