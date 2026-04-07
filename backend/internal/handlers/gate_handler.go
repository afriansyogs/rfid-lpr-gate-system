package handlers

import (
	"github.com/gofiber/fiber/v3"
	"rfid-gateway/internal/dto"
	"fmt"
)

func HandleRfidTap(c fiber.Ctx) error {
	var req dto.TapRequest

	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: "Format JSON not valid",
		})
	}

	fmt.Printf(">>> RFID Tap: [%s] from ESP32\n", req.RfidUUID)

	return c.JSON(dto.TapResponse{
		Status:    "success",
		Action:    "hold",
		RfidUUID: req.RfidUUID,
		Message:   "Data received, waiting for AI logic...",
	})
}