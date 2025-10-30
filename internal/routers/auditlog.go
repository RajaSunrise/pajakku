package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteAuditLog(app *fiber.App) {
	// Initialize dependencies
	auditLogRepo := repository.NewAuditLogRepository(databases.DB)
	svc := service.NewAuditLogService(auditLogRepo)
	handler := handlers.NewAuditLogHandler(svc)

	// Routes
	auditLogs := app.Group("/api/v1/auditlogs")
	auditLogs.Post("/", middlewares.JWTAuth, handler.CreateAuditLog)
	auditLogs.Get("/:id", middlewares.JWTAuth, handler.GetAuditLogByID)
	auditLogs.Get("/user", middlewares.JWTAuth, handler.GetAuditLogsByUserID)
	auditLogs.Get("/", middlewares.JWTAuth, handler.GetAllAuditLogs)
}
