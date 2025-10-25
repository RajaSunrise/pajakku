package handlers

import (
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
)


type UsersProfileHandler struct {
	service  service.UserProfileService
}


func NewUsersProfileHandler(service service.UserProfileService) *UsersProfileHandler {
	return &UsersProfileHandler{service: service}
}


func (h *UsersProfileHandler) CreateUsersProfile(c *fiber.Ctx) error {
	var req request.CreateUsersProfile
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}
	resp, err := h.service.CreateProfile(&req);
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}


func (h *UsersProfileHandler) GetProfileByID(c *fiber.Ctx) error {
	id,_ := strconv.Atoi(c.Params("id"))

	resp, err := h.service.GetProfileByID(uint(id));
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}


func (h *UsersProfileHandler) UpdateUsersProfile(c *fiber.Ctx) error {
	id, _:= strconv.Atoi(c.Params("id"))
	var req request.UpdateUsersProfile
	
	resp, err :=  h.service.UpdateProfile(uint(id), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}