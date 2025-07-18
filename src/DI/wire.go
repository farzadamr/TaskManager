package DI

import (
	"log"

	"github.com/farzadamr/TaskManager/api/handlers"
	"github.com/farzadamr/TaskManager/config"
	database "github.com/farzadamr/TaskManager/db"
	"github.com/farzadamr/TaskManager/models"
	"github.com/farzadamr/TaskManager/repositories"
	"github.com/farzadamr/TaskManager/services"
	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	TaskHandler *handlers.TaskHandler
	AuthHandler *handlers.AuthHandler
}

func NewContainer() (*Container, error) {
	db, err := database.InitDB(config.LoadConfig())
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
		return nil, err
	}
	//MigrateModels
	if err := database.MigrateModels(db, &models.Task{}); err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	//Initialize Repositories
	taskRepo := repositories.NewTaskRepository(db)
	userRepo := repositories.NewUserRepository(db)

	//Initialize Services
	taskService := services.NewTaskService(taskRepo)
	authService := services.NewAuthService(userRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	authHandler := handlers.NewAuthHandler(authService)

	return &Container{
		DB:          db,
		TaskHandler: taskHandler,
		AuthHandler: authHandler,
	}, nil
}
func (c *Container) Close() {
	database.CloseDB(c.DB)
}
