package routes

import (
	"rfid-gateway/internal/handlers"
	"rfid-gateway/internal/repositories"
	"rfid-gateway/internal/services"
	"github.com/gofiber/fiber/v3"
)

func SetupMemberRoutes(api fiber.Router, repo *repositories.SupabaseRepo) {
	memberGroup := api.Group("/members")
	
	memberService := &services.MemberService{
		Repo: repo,
	}

	memberHandler := &handlers.MemberHandler{
		Service: memberService,
	}

	memberGroup.Get("", memberHandler.GetMembers)
	memberGroup.Post("", memberHandler.CreateMember)
	memberGroup.Delete("/:id", memberHandler.DeleteMember)
	memberGroup.Patch("/:id", memberHandler.UpdateMember)
}