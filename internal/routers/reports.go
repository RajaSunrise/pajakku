package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteReport(app *fiber.App) {
	// Initialize dependencies
	reportRepo := repository.NewReportRepository(databases.DB)
	profileRepo := repository.NewUserProfileRepository(databases.DB)
	svc := service.NewReportService(reportRepo, profileRepo)
	handler := handlers.NewReportHandler(svc)

	// Routes
	reports := app.Group("/api/v1/reports")
	reports.Post("/", middlewares.JWTAuth, handler.CreateReport)
	reports.Get("/:id", middlewares.JWTAuth, handler.GetReportByID)
	reports.Get("/", middlewares.JWTAuth, handler.GetReportsByUserID)
	reports.Put("/:id", middlewares.JWTAuth, handler.UpdateReport)
	reports.Delete("/:id", middlewares.JWTAuth, handler.DeleteReport)
}
