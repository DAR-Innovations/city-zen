package main

import (
	"fmt"
	"github.com/DAR-Innovations/city-zen/internal/config"
	"github.com/DAR-Innovations/city-zen/internal/database"
	"github.com/DAR-Innovations/city-zen/internal/modules/auth"
	"github.com/DAR-Innovations/city-zen/internal/routes"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	// Load configuration from the .env file
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Fiber
	app := fiber.New()

	// Dynamic Routes
	api := app.Group("/api/v1")
	// Configure routes dynamically here (e.g., using a router function)

	DBHandler := InitializeDatabase(config.GetConfig())

	authService := auth.NewAuthenticationService(DBHandler.DB)
	authHandler := auth.NewAuthenticationHandler(authService)

	routes.RegisterAuthRoutes(api, *authHandler)

	// Start the server
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
