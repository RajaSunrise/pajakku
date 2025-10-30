package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteRole(app *fiber.App) {
	// Initialize dependencies
	roleRepo := repository.NewRoleRepository(databases.DB)
	svc := service.NewRoleService(roleRepo)
	handler := handlers.NewRoleHandler(svc)

	// Routes
	roles := app.Group("/api/v1/roles")
	roles.Post("/", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.CreateRole)
	roles.Get("/:id", middlewares.JWTAuth, handler.GetRoleByID)
	roles.Get("/name/:name", middlewares.JWTAuth, handler.GetRoleByName)
	roles.Get("/", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.GetAllRoles)
	roles.Put("/:id", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.UpdateRole)
	roles.Delete("/:id", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.DeleteRole)
}
