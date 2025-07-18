package main

import (
	"log"

	"github.com/farzadamr/TaskManager/DI"
	"github.com/farzadamr/TaskManager/api"
	"github.com/farzadamr/TaskManager/api/validators"
	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Container
	container, err := DI.NewContainer()
	if err != nil {
		log.Fatalf("failed to initialize container: %v", err)
	}
	defer container.Close()

	e := echo.New()
	e.Validator = validators.NewValidator()
	api.SetupRoutes(
		e,
		container.AuthHandler,
		container.TaskHandler,
	)

	port := ":8080"
	log.Printf("Server started on port %s", port)
	if err := e.Start(port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
