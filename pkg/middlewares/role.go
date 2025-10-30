package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func RoleAuth(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleID := c.Locals("roleID").(uint)
		// Assuming roleID 1 is admin, 2 is user, etc.
		// You can fetch role name from DB or hardcode
		if requiredRole == "admin" && roleID != 1 {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}
		// For user, allow if roleID == 2 or something
		return c.Next()
	}
}
