package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BillingHandler struct {
	service service.BillingService
}

func NewBillingHandler(service service.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) CreateBilling(c *fiber.Ctx) error {
	logrus.Info("Create billing request received")
	userID := c.Locals("userID").(string)

	var req request.CreateBillingRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create billing request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateBilling(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create billing")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Billing created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *BillingHandler) GetBillingByID(c *fiber.Ctx) error {
	logrus.Info("Get billing by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetBillingByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Billing not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Billing not found"})
	}

	logrus.WithField("billingID", idStr).Info("Billing retrieved successfully")
	return c.JSON(resp)
}

func (h *BillingHandler) GetBillingsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get billings by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetBillingsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get billings")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Billings retrieved successfully")
	return c.JSON(resp)
}

func (h *BillingHandler) UpdateBilling(c *fiber.Ctx) error {
	logrus.Info("Update billing request received")
	idStr := c.Params("id")

	var req request.UpdateBillingRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update billing request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateBilling(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update billing")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("billingID", idStr).Info("Billing updated successfully")
	return c.JSON(resp)
}

func (h *BillingHandler) DeleteBilling(c *fiber.Ctx) error {
	logrus.Info("Delete billing request received")
	idStr := c.Params("id")

	err := h.service.DeleteBilling(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete billing")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("billingID", idStr).Info("Billing deleted successfully")
	return c.JSON(fiber.Map{"message": "Billing deleted successfully"})
}
