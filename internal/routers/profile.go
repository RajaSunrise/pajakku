package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteProfile(app *fiber.App) {
	// Initialization Depencies
	repo := repository.NewUserProfileRepository(databases.DB)
	svc := service.NewUserProfileService(repo)
	handler := handlers.NewUsersProfileHandler(svc)

	// Routes Data Pajak/Profile Pajak
	users := app.Group("/api/v1/users")
	users.Post("/", middlewares.JWTAuth(), handler.CreateUsersProfile)
	users.Post("/:id", middlewares.JWTAuth(), handler.UpdateUsersProfile)
	users.Get("/:id", middlewares.JWTAuth(), handler.GetProfileByID)
}