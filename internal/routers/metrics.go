package routers

import (
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteMetrics(app *fiber.App) {
	handler := handlers.NewMetricsHandler()

	// Metrics endpoint for Prometheus
	app.Get("/metrics", handler.GetMetrics)
}
