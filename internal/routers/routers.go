package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Setup Middlewares
	SetupMiddleware(app)

	// Setup routes auth
	SetupRouteAuth(app)

	// Setup routes users
	SetupRouteUser(app)

	// Setup routes roles
	SetupRouteRole(app)

	// Setup routes attachments
	SetupRouteAttachment(app)

	// Setup routes audit logs
	SetupRouteAuditLog(app)

	// Setup routes invoices
	SetupRouteInvoice(app)

	// Setup routes notifications
	SetupRouteNotification(app)

	// Setup routes payments
	SetupRoutePayment(app)

	// Setup routes tax types
	SetupRouteTaxType(app)

	// Setup routes metrics
	SetupRouteMetrics(app)
}
