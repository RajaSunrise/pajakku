package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteAuth(app *fiber.App) {
	// Initialize dependencies
	repo := repository.NewUsersAuthRepository(databases.DB)
	svc := service.NewUsersAuthService(repo)
	handler := handlers.NewUsersAuthHandler(svc)

	// Routes
	auth := app.Group("/")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/forget-password", handler.ForgetPassword)
	auth.Post("/reset-password", handler.ResetPassword)
	auth.Post("/logout", handler.Logout)
}
