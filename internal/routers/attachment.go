package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteAttachment(app *fiber.App) {
	// Initialize dependencies
	attachmentRepo := repository.NewAttachmentRepository(databases.DB)
	svc := service.NewAttachmentService(attachmentRepo)
	handler := handlers.NewAttachmentHandler(svc)

	// Routes
	attachments := app.Group("/api/v1/attachments")
	attachments.Post("/", middlewares.JWTAuth, handler.CreateAttachment)
	attachments.Get("/:id", middlewares.JWTAuth, handler.GetAttachmentByID)
	attachments.Get("/", middlewares.JWTAuth, handler.GetAttachmentsByUserID)
	attachments.Put("/:id", middlewares.JWTAuth, handler.UpdateAttachment)
	attachments.Delete("/:id", middlewares.JWTAuth, handler.DeleteAttachment)
}
