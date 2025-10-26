package routers

import (
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupMiddleware(app *fiber.App)  {
	app.Use(middlewares.CORS())
	app.Use(middlewares.RateLimiter())
}