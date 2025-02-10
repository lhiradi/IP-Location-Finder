package main

import (
	"ip-api/config"
	"ip-api/internal/database"
	"ip-api/internal/handlers"
	"ip-api/internal/middlewares"
	"ip-api/internal/repositories"
	"ip-api/internal/routes"
	"ip-api/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LoadConfig()

	// Initialize the database connection
	db, err := database.InitDB(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	ipRepo := repositories.NewIPRepo(db)
	ipService := services.NewIPServ(ipRepo)
	ipHandler := handlers.NewIPHandler(ipService)

	app := fiber.New()

	// Setup routes
	api := routes.SetupRoutes(app)
	api.Use(middlewares.GetIPMiddleware(ipService))
	api.Post("/ip", ipHandler.IPHandler)

	err = app.Listen(":8080")
	log.Println("Listening on port 8080: ")

	if err != nil {
		log.Fatal(err)
	}
}
