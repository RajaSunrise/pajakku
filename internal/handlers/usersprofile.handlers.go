package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UsersProfileHandler struct {
	service service.UserProfileService
}

func NewUsersProfileHandler(service service.UserProfileService) *UsersProfileHandler {
	return &UsersProfileHandler{service: service}
}

func (h *UsersProfileHandler) CreateUsersProfile(c *fiber.Ctx) error {
	logrus.WithField("userID", c.Locals("userID")).Info("Create profile request received")
	userID := c.Locals("userID").(string)
	var req request.CreateUsersProfile
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).WithField("userID", userID).Warn("Failed to parse create profile request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}
	resp, err := h.service.CreateProfile(userID, &req)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create profile")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.WithField("userID", userID).Info("Profile created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UsersProfileHandler) GetProfileByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	logrus.WithField("userID", userID).Info("Get profile request received")

	resp, err := h.service.GetProfileByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Warn("Profile not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Profile not found",
		})
	}

	logrus.WithField("userID", userID).Info("Profile retrieved successfully")
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *UsersProfileHandler) UpdateUsersProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	logrus.WithField("userID", userID).Info("Update profile request received")
	var req request.UpdateUsersProfile
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).WithField("userID", userID).Warn("Failed to parse update profile request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	// Get profile by userID first
	profileResp, err := h.service.GetProfileByUserID(userID)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Warn("Profile not found for update")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Profile not found",
		})
	}

	resp, err := h.service.UpdateProfileByNIK(profileResp.NIK, &req)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to update profile")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	logrus.WithField("userID", userID).Info("Profile updated successfully")
	return c.Status(fiber.StatusOK).JSON(resp)
}
