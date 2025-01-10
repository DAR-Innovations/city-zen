package main

import (
	"fmt"
	"log"

	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/DAR-Innovations/city-zen/internal/database"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth"
	"github.com/DAR-Innovations/city-zen/internal/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app := fiber.New()

	api := app.Group("/api/v1")

	DBHandler := InitializeDatabase(config.GetConfig())

	authService := auth.NewAuthenticationService(DBHandler.DB)
	authHandler := auth.NewAuthenticationHandler(authService)

	routes.RegisterAuthRoutes(api, *authHandler)
	routes.RegisterImagesRoutes(api)

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Starting server at %s", address)
	if err := app.Listen(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}

func InitializeDatabase(cfg *config.Config) *database.DBHandler {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
	)

	dbHandler, err := database.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	return dbHandler
}
