package middlewares

import (
	"ip-api/internal/models"
	"ip-api/internal/services"

	"github.com/gofiber/fiber/v2"
)

// GetIPMiddleware checks if the IP exists in the database
func GetIPMiddleware(ipService services.IPServ) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.Query("ip")
		if ip == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "IP required!"})
		}

		var ipData models.ResponseIP
		if retrievdIP, err := ipService.GetIPService(ip); err == nil {

			ipData.IP = retrievdIP.IP
			ipData.CountryName = retrievdIP.CountryName
			return c.JSON(&ipData)
		}

		return c.Next()
	}
}
