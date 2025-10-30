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
	userRepo := repository.NewUserRepository(databases.DB)
	roleRepo := repository.NewRoleRepository(databases.DB)
	userSvc := service.NewUserService(userRepo, roleRepo)
	authHandler := handlers.NewAuthHandler(userSvc)

	// Routes
	auth := app.Group("/api/v1/auth")
	auth.Post("/signup", authHandler.Signup)
	auth.Post("/login", authHandler.Login)
}
