package handlers

import (
	"strings"

	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UsersAuthHandler struct {
	service service.UsersAuthService
}

func NewUsersAuthHandler(service service.UsersAuthService) *UsersAuthHandler {
	return &UsersAuthHandler{service: service}
}

func (h *UsersAuthHandler) Register(c *fiber.Ctx) error {
	logrus.Info("Register request received")
	var req request.CreateUsersAuth
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse register request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.Register(&req); err != nil {
		logrus.WithError(err).Error("Failed to register user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("email", req.Email).Info("User registered successfully")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *UsersAuthHandler) Login(c *fiber.Ctx) error {
	logrus.Info("Login request received")
	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse login request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.Login(&req)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Warn("Login failed")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("email", req.Email).Info("User logged in successfully")
	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *UsersAuthHandler) ForgetPassword(c *fiber.Ctx) error {
	logrus.Info("Forget password request received")
	var req request.ForgetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse forget password request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.ForgetPassword(&req); err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to process forget password")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("email", req.Email).Info("Password reset email sent")
	return c.Status(fiber.StatusOK).JSON(response.PasswordResetResponse{Message: "Password reset email sent"})
}

func (h *UsersAuthHandler) ResetPassword(c *fiber.Ctx) error {
	logrus.Info("Reset password request received")
	var req request.ResetPasswordRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse reset password request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.ResetPassword(&req); err != nil {
		logrus.WithError(err).Error("Failed to reset password")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("Password reset successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Password reset successfully"})
}

func (h *UsersAuthHandler) Logout(c *fiber.Ctx) error {
	logrus.Info("Logout request received")
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		logrus.Warn("Missing authorization header in logout request")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		logrus.Warn("Invalid token format in logout request")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}

	err := utils.InvalidateToken(tokenString)
	if err != nil {
		logrus.WithError(err).Error("Failed to invalidate token during logout")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to logout"})
	}

	logrus.Info("User logged out successfully")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logged out successfully"})
}
