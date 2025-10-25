package routers

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App)  {
	// Setup routes Auth
	SetupRouteAuth(app)

	// Setup routes Users
	SetupRouteProfile(app)

	// Setup routes reports
	SetupRouteReport(app)

	// Setup routes billings
	SetupRouteBiling(app)
}