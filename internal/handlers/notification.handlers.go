package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NotificationHandler struct {
	service service.NotificationService
}

func NewNotificationHandler(service service.NotificationService) *NotificationHandler {
	return &NotificationHandler{service: service}
}

func (h *NotificationHandler) CreateNotification(c *fiber.Ctx) error {
	logrus.Info("Create notification request received")
	userID := c.Locals("userID").(string)

	var req request.CreateNotification
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create notification request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateNotification(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create notification")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Notification created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *NotificationHandler) GetNotificationByID(c *fiber.Ctx) error {
	logrus.Info("Get notification by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetNotificationByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Notification not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Notification not found"})
	}

	logrus.WithField("notificationID", idStr).Info("Notification retrieved successfully")
	return c.JSON(resp)
}

func (h *NotificationHandler) GetNotificationsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get notifications by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetNotificationsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get notifications")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Notifications retrieved successfully")
	return c.JSON(resp)
}

func (h *NotificationHandler) UpdateNotification(c *fiber.Ctx) error {
	logrus.Info("Update notification request received")
	idStr := c.Params("id")

	var req request.UpdateNotification
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update notification request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateNotification(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update notification")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("notificationID", idStr).Info("Notification updated successfully")
	return c.JSON(resp)
}

func (h *NotificationHandler) DeleteNotification(c *fiber.Ctx) error {
	logrus.Info("Delete notification request received")
	idStr := c.Params("id")

	err := h.service.DeleteNotification(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete notification")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("notificationID", idStr).Info("Notification deleted successfully")
	return c.JSON(fiber.Map{"message": "Notification deleted successfully"})
}
