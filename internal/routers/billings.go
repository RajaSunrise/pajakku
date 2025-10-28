package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteBiling(app *fiber.App) {
	// Initialize dependencies
	billingRepo := repository.NewBillingRepository(databases.DB)
	profileRepo := repository.NewUserProfileRepository(databases.DB)
	svc := service.NewBillingService(billingRepo, profileRepo)
	handler := handlers.NewBillingHandler(svc)

	// Routes
	billings := app.Group("/api/v1/billings")
	billings.Post("/", middlewares.JWTAuth, handler.CreateBilling)
	billings.Get("/:id", middlewares.JWTAuth, handler.GetBillingByID)
	billings.Get("/", middlewares.JWTAuth, handler.GetBillingsByUserID)
	billings.Put("/:id", middlewares.JWTAuth, handler.UpdateBilling)
	billings.Delete("/:id", middlewares.JWTAuth, handler.DeleteBilling)
}
