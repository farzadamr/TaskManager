package db

import (
	"fmt"
	"log"

	"github.com/farzadamr/TaskManager/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	var dialector gorm.Dialector

	switch cfg.DBDriver {
	case "sqlite":
		dialector = sqlite.Open(cfg.DBName)
	// case "postgres":
	// 	dsn := fmt.Sprintf("host:%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	// 	dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.DBDriver)
	}
	DB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed connect to database")
	}
	log.Println("Database connection established")
	return DB, nil
}
func MigrateModels(db *gorm.DB, models ...interface{}) error {
	if err := db.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}
	log.Println("Database migration completed")
	return nil
}

func ColseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("failed to get database instance: %v", err)
		return
	}
	if err = sqlDB.Close(); err != nil {
		log.Printf("failed to close database connection: %v", err)
	}
	log.Println("Database connection closed")
}
