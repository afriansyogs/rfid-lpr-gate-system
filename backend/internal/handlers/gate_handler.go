package handlers

import (
	"github.com/gofiber/fiber/v3"
	"rfid-gateway/internal/dto"
	"rfid-gateway/internal/services"
)

type GateHandler struct {
	GateService *services.GateService
}

func NewGateHandler(gateService *services.GateService) *GateHandler {
	return &GateHandler{GateService: gateService}
}

func (h *GateHandler) HandleTap(c fiber.Ctx) error {
	var req dto.TapRequest

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Status: "ERROR",
			Reason: "Format JSON tidak valid",
		})
	}

	response := h.GateService.ProcessTap(req)

	if response.Status == "DENIED" {
		return c.Status(fiber.StatusForbidden).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}