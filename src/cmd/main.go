package main

import (
	"log"

	"github.com/farzadamr/TaskManager/api"
	"github.com/farzadamr/TaskManager/api/validators"
	"github.com/farzadamr/TaskManager/config"
	database "github.com/farzadamr/TaskManager/db"
	"github.com/farzadamr/TaskManager/models"
	"github.com/farzadamr/TaskManager/repositories"
	"github.com/farzadamr/TaskManager/services"
	"github.com/labstack/echo/v4"
)

func main() {
	//Load Config
	cfg := config.LoadConfig()

	// Initialize Database
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.CloseDB(db)

	//MigrateModels
	if err := database.MigrateModels(db, &models.Task{}); err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)

	e := echo.New()
	e.Validator = validators.NewValidator()
	api.SetupRoutes(e, taskService)

	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
