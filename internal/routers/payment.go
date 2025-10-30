package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutePayment(app *fiber.App) {
	// Initialize dependencies
	paymentRepo := repository.NewPaymentRepository(databases.DB)
	svc := service.NewPaymentService(paymentRepo)
	handler := handlers.NewPaymentHandler(svc)

	// Routes
	payments := app.Group("/api/v1/payments")
	payments.Post("/", middlewares.JWTAuth, handler.CreatePayment)
	payments.Get("/:id", middlewares.JWTAuth, handler.GetPaymentByID)
	payments.Get("/", middlewares.JWTAuth, handler.GetPaymentsByUserID)
	payments.Put("/:id", middlewares.JWTAuth, handler.UpdatePayment)
	payments.Delete("/:id", middlewares.JWTAuth, handler.DeletePayment)
}
