package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the application
func SetupRoutes(app *fiber.App) fiber.Router {
	api := app.Group("/api/v1")
	return api
}
