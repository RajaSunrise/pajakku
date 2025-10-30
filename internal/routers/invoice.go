package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteInvoice(app *fiber.App) {
	// Initialize dependencies
	invoiceRepo := repository.NewInvoiceRepository(databases.DB)
	svc := service.NewInvoiceService(invoiceRepo)
	handler := handlers.NewInvoiceHandler(svc)

	// Routes
	invoices := app.Group("/api/v1/invoices")
	invoices.Post("/", middlewares.JWTAuth, handler.CreateInvoice)
	invoices.Get("/:id", middlewares.JWTAuth, handler.GetInvoiceByID)
	invoices.Get("/", middlewares.JWTAuth, handler.GetInvoicesByUserID)
	invoices.Put("/:id", middlewares.JWTAuth, handler.UpdateInvoice)
	invoices.Delete("/:id", middlewares.JWTAuth, handler.DeleteInvoice)
}
