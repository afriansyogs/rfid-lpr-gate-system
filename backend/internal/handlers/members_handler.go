package handlers

import (
	"strconv"
	"rfid-gateway/internal/dto"
	"rfid-gateway/internal/services"
	"github.com/gofiber/fiber/v3"
)

var members []dto.Member

type MemberHandler struct {
	Service *services.MemberService
}

func (h *MemberHandler) GetMembers(c fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	limitStr := c.Query("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	
	name := c.Query("name")
	rfid := c.Query("rfid")
	uhf := c.Query("uhf")
	plate := c.Query("plate")
	telegramChatID := c.Query("telegram_chat_id")

	results, err := h.Service.FetchMembers(page, limit, name, rfid, uhf, plate, telegramChatID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}

	if len(results.Data) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: "Data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func (h *MemberHandler) CreateMember(c fiber.Ctx) error {
	var newMember dto.Member

	if err := c.Bind().Body(&newMember); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}

	err := h.Service.CreateMember(&newMember)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(newMember)
}

func (h *MemberHandler) DeleteMember(c fiber.Ctx) error {
	id := c.Params("id")

	err := h.Service.DeleteMember(id); 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *MemberHandler) UpdateMember(c fiber.Ctx) error {
	id := c.Params("id")
	var updatedData dto.Member
	if err := c.Bind().Body(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}

	err := h.Service.UpdateMember(id, &updatedData)
	
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status: "error",
			Reason: err.Error(),
		})
	}
	
	updatedData.ID = id

	return c.Status(fiber.StatusOK).JSON(updatedData)
}