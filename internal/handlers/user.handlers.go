package handlers

import (
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	logrus.Info("Create user request received")

	var req request.CreateUser
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create user request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateUser(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("User created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	logrus.Info("Get user by ID request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	resp, err := h.service.GetUserByID(uint(id))
	if err != nil {
		logrus.WithError(err).Warn("User not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	logrus.WithField("userID", id).Info("User retrieved successfully")
	return c.JSON(resp)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	logrus.Info("Get all users request received")

	userID := c.Locals("userID").(uint)
	roleID := c.Locals("roleID").(uint)

	if roleID != 1 {
		// Not admin, return only current user's data
		resp, err := h.service.GetUserByID(userID)
		if err != nil {
			logrus.WithError(err).Warn("User not found")
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		logrus.Info("Current user data retrieved successfully")
		return c.JSON([]response.UserResponse{*resp})
	}

	// Admin, return all users
	resp, err := h.service.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all users")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.Info("All users retrieved successfully")
	return c.JSON(resp)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	logrus.Info("Update user request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userID := c.Locals("userID").(uint)
	roleID := c.Locals("roleID").(uint)

	var req request.UpdateUser
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update user request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if roleID != 1 {
		// Not admin
		if uint(id) != userID {
			logrus.Warn("Non-admin user trying to update another user")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You can only update your own data"})
		}
		// Prevent non-admin from updating role and status
		if req.RoleID != 0 || req.StatusAktif != nil {
			logrus.Warn("Non-admin user trying to update restricted fields")
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You cannot update role or status"})
		}
	}

	resp, err := h.service.UpdateUser(uint(id), &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", id).Info("User updated successfully")
	return c.JSON(resp)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	logrus.Info("Delete user request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid user ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	err = h.service.DeleteUser(uint(id))
	if err != nil {
		logrus.WithError(err).Error("Failed to delete user")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", id).Info("User deleted successfully")
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
