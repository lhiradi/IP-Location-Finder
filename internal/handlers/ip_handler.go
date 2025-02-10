package handlers

import (
	"encoding/json"
	"ip-api/internal/external"
	"ip-api/internal/models"
	"ip-api/internal/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type IPHandler struct {
	ipServ services.IPServ
}

func NewIPHandler(ipServ services.IPServ) IPHandler {
	return IPHandler{ipServ: ipServ}
}

// IPHandler handles the request for IP data
func (h *IPHandler) IPHandler(c *fiber.Ctx) error {
// 8080/api/v1/ip?=
	ip := c.Query("ip")
	if ip == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "IP required!"})
	}

	// Call the external API
	resp, err := external.CallExternalAPI(ip)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve data"})
	}
	defer resp.Body.Close()

	// Decode the response
	var apiResponse models.IP
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode gathered data"})
	}

	if err := h.ipServ.SaveIPService(apiResponse); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Couldn't save ip"})
	}

	var ipResponse models.ResponseIP
	ipResponse.IP = apiResponse.IP
	ipResponse.CountryName = apiResponse.CountryName

	return c.JSON(&ipResponse)
}
