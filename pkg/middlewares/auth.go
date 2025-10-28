package middlewares

import (
	"strings"

	"github.com/RajaSunrise/pajakku/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func JWTAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Token"})
	}
	
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
	}
	
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
	}
	
	c.Locals("userID", claims.UserID)
	c.Locals("email", claims.Email)
	
	return c.Next()
}