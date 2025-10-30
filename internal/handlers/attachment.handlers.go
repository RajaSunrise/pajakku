package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AttachmentHandler struct {
	service service.AttachmentService
}

func NewAttachmentHandler(service service.AttachmentService) *AttachmentHandler {
	return &AttachmentHandler{service: service}
}

func (h *AttachmentHandler) CreateAttachment(c *fiber.Ctx) error {
	logrus.Info("Create attachment request received")
	userID := c.Locals("userID").(string)

	var req request.CreateAttachment
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create attachment request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateAttachment(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create attachment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Attachment created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AttachmentHandler) GetAttachmentByID(c *fiber.Ctx) error {
	logrus.Info("Get attachment by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetAttachmentByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Attachment not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Attachment not found"})
	}

	logrus.WithField("attachmentID", idStr).Info("Attachment retrieved successfully")
	return c.JSON(resp)
}

func (h *AttachmentHandler) GetAttachmentsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get attachments by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetAttachmentsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get attachments")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Attachments retrieved successfully")
	return c.JSON(resp)
}

func (h *AttachmentHandler) UpdateAttachment(c *fiber.Ctx) error {
	logrus.Info("Update attachment request received")
	idStr := c.Params("id")

	var req request.UpdateAttachment
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update attachment request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateAttachment(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update attachment")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("attachmentID", idStr).Info("Attachment updated successfully")
	return c.JSON(resp)
}

func (h *AttachmentHandler) DeleteAttachment(c *fiber.Ctx) error {
	logrus.Info("Delete attachment request received")
	idStr := c.Params("id")

	err := h.service.DeleteAttachment(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete attachment")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("attachmentID", idStr).Info("Attachment deleted successfully")
	return c.JSON(fiber.Map{"message": "Attachment deleted successfully"})
}
