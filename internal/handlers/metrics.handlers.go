package handlers

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct{}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{}
}

func (h *MetricsHandler) GetMetrics(c *fiber.Ctx) error {
	// Use adaptor to convert promhttp.Handler to Fiber handler
	return adaptor.HTTPHandler(promhttp.Handler())(c)
}
