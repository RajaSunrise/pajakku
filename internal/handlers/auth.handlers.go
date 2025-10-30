package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
	logrus.Info("Signup request received")

	var signupReq request.SignupRequest
	if err := c.BodyParser(&signupReq); err != nil {
		logrus.WithError(err).Warn("Failed to parse signup request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Create CreateUser request with default role
	req := request.CreateUser{
		NIK:             signupReq.NIK,
		NPWP:            signupReq.NPWP,
		Nama:            signupReq.Nama,
		Email:           signupReq.Email,
		Password:        signupReq.Password,
		Alamat:          signupReq.Alamat,
		JenisWajibPajak: signupReq.JenisWajibPajak,
		RoleID:          2, // Default role ID for user
	}

	resp, err := h.userService.CreateUser(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to signup user")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("User signed up successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	logrus.Info("Login request received")

	var req request.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse login request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.userService.Login(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to login user")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("email", req.Email).Info("User logged in successfully")
	return c.JSON(resp)
}
