package routers

import (
	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/handlers"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/RajaSunrise/pajakku/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRouteTaxType(app *fiber.App) {
	// Initialize dependencies
	taxTypeRepo := repository.NewTaxTypeRepository(databases.DB)
	svc := service.NewTaxTypeService(taxTypeRepo)
	handler := handlers.NewTaxTypeHandler(svc)

	// Routes
	taxTypes := app.Group("/api/v1/taxtypes")
	taxTypes.Post("/", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.CreateTaxType)
	taxTypes.Get("/:id", middlewares.JWTAuth, handler.GetTaxTypeByID)
	taxTypes.Get("/code/:code", middlewares.JWTAuth, handler.GetTaxTypeByCode)
	taxTypes.Get("/", middlewares.JWTAuth, handler.GetAllTaxTypes)
	taxTypes.Put("/:id", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.UpdateTaxType)
	taxTypes.Delete("/:id", middlewares.JWTAuth, middlewares.RoleAuth("admin"), handler.DeleteTaxType)
}
