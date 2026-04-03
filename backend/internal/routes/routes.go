package routes

import (
	"rfid-gateway/internal/dto"
	"github.com/gofiber/fiber/v3"
	"rfid-gateway/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(dto.HealthResponse{
			Status:  "ok",
			Message: "Gateway Sistem RFID LPR ready",
		})
	})

	api.Post("/gate/tap", handlers.HandleRfidTap)
}