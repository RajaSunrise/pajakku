package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type PaymentHandler struct {
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (h *PaymentHandler) CreatePayment(c *fiber.Ctx) error {
	logrus.Info("Create payment request received")
	userID := c.Locals("userID").(string)

	var req request.CreatePayment
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create payment request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreatePayment(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create payment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Payment created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *PaymentHandler) GetPaymentByID(c *fiber.Ctx) error {
	logrus.Info("Get payment by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetPaymentByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Payment not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment not found"})
	}

	logrus.WithField("paymentID", idStr).Info("Payment retrieved successfully")
	return c.JSON(resp)
}

func (h *PaymentHandler) GetPaymentsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get payments by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetPaymentsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get payments")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Payments retrieved successfully")
	return c.JSON(resp)
}

func (h *PaymentHandler) UpdatePayment(c *fiber.Ctx) error {
	logrus.Info("Update payment request received")
	idStr := c.Params("id")

	var req request.UpdatePayment
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update payment request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdatePayment(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update payment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("paymentID", idStr).Info("Payment updated successfully")
	return c.JSON(resp)
}

func (h *PaymentHandler) DeletePayment(c *fiber.Ctx) error {
	logrus.Info("Delete payment request received")
	idStr := c.Params("id")

	err := h.service.DeletePayment(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete payment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("paymentID", idStr).Info("Payment deleted successfully")
	return c.JSON(fiber.Map{"message": "Payment deleted successfully"})
}
