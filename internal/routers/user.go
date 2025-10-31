package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteUser(app *fiber.App) {
	// Initialize dependencies
	userRepo := repository.NewUserRepository(databases.DB)
	roleRepo := repository.NewRoleRepository(databases.DB)
	svc := service.NewUserService(userRepo, roleRepo)
	handler := handlers.NewUserHandler(svc)

	// Routes
	users := app.Group("/api/v1/users")
	users.Post("/", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.CreateUser)
	users.Get("/:id", middlewares.JWTAuth, handler.GetUserByID)
	users.Get("/", middlewares.JWTAuth, handler.GetAllUsers)
	users.Put("/:id", middlewares.JWTAuth, handler.UpdateUser)
	users.Delete("/:id", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.DeleteUser)
}
