package main

import (
	"log"

	"github.com/farzadamr/TaskManager/config"
	database "github.com/farzadamr/TaskManager/db"
	"github.com/farzadamr/TaskManager/models"
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
}
