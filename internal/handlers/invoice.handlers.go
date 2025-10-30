package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type InvoiceHandler struct {
	service service.InvoiceService
}

func NewInvoiceHandler(service service.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

func (h *InvoiceHandler) CreateInvoice(c *fiber.Ctx) error {
	logrus.Info("Create invoice request received")
	userID := c.Locals("userID").(string)

	var req request.CreateInvoice
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create invoice request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateInvoice(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create invoice")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Invoice created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *InvoiceHandler) GetInvoiceByID(c *fiber.Ctx) error {
	logrus.Info("Get invoice by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetInvoiceByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Invoice not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invoice not found"})
	}

	logrus.WithField("invoiceID", idStr).Info("Invoice retrieved successfully")
	return c.JSON(resp)
}

func (h *InvoiceHandler) GetInvoicesByUserID(c *fiber.Ctx) error {
	logrus.Info("Get invoices by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetInvoicesByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get invoices")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Invoices retrieved successfully")
	return c.JSON(resp)
}

func (h *InvoiceHandler) UpdateInvoice(c *fiber.Ctx) error {
	logrus.Info("Update invoice request received")
	idStr := c.Params("id")

	var req request.UpdateInvoice
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update invoice request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateInvoice(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update invoice")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("invoiceID", idStr).Info("Invoice updated successfully")
	return c.JSON(resp)
}

func (h *InvoiceHandler) DeleteInvoice(c *fiber.Ctx) error {
	logrus.Info("Delete invoice request received")
	idStr := c.Params("id")

	err := h.service.DeleteInvoice(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete invoice")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("invoiceID", idStr).Info("Invoice deleted successfully")
	return c.JSON(fiber.Map{"message": "Invoice deleted successfully"})
}
