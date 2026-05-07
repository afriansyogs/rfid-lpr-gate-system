package routes

import (
	"rfid-gateway/internal/dto"
	"rfid-gateway/internal/handlers"
	"rfid-gateway/internal/services"
	"rfid-gateway/internal/repositories"
	"github.com/gofiber/fiber/v3"
	"github.com/nedpals/supabase-go"
)

func SetupRoutes(app *fiber.App, db *supabase.Client, aiServiceURL string) {
	supabaseRepo := repositories.NewSupabaseRepo(db)
	aiRepo := repositories.NewAIRepo(aiServiceURL)

	gateService := services.NewGateService(supabaseRepo, aiRepo)
	gateHandler := handlers.NewGateHandler(gateService)

	api := app.Group("/api")

	api.Get("/health", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(dto.HealthResponse{
			Status:  "OK",
			Message: "Gate Backend Service is running smoothly",
		})
	})

	api.Post("/gate/tap", gateHandler.HandleTap)

	SetupMemberRoutes(api, supabaseRepo)
	SetupLogsRoutes(api, supabaseRepo)
}