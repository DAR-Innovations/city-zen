package database

import (
	"fmt"
	"log"
	"time"

	"github.com/DAR-Innovations/city-zen/internal/data"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func InitDB(dsn string) (*DBHandler, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	/*if err := applyMigrations(db, "migrations"); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}*/

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get SQL DB from gorm DB: %w", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected and migrations applied successfully.")
	return &DBHandler{DB: db}, nil
}

func applyAutoMigrate(db *gorm.DB) error {
	log.Println("Starting AutoMigrate...")

	models := []interface{}{
		&data.Department{},
		&data.User{},
		&data.TaskType{},
		&data.DepartmentTaskType{},
		&data.Issue{},
		&data.VolunteerTask{},
		&data.DepartmentTask{},
		&data.UserReport{},
		&data.DepartmentReport{},
		&data.Employee{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate model %T: %w", model, err)
		}
	}

	log.Println("AutoMigrate completed successfully.")
	return nil
}
