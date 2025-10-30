package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteNotification(app *fiber.App) {
	// Initialize dependencies
	notificationRepo := repository.NewNotificationRepository(databases.DB)
	svc := service.NewNotificationService(notificationRepo)
	handler := handlers.NewNotificationHandler(svc)

	// Routes
	notifications := app.Group("/api/v1/notifications")
	notifications.Post("/", middlewares.JWTAuth, handler.CreateNotification)
	notifications.Get("/:id", middlewares.JWTAuth, handler.GetNotificationByID)
	notifications.Get("/", middlewares.JWTAuth, handler.GetNotificationsByUserID)
	notifications.Put("/:id", middlewares.JWTAuth, handler.UpdateNotification)
	notifications.Delete("/:id", middlewares.JWTAuth, handler.DeleteNotification)
}
