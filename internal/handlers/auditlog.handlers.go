package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuditLogHandler struct {
	service service.AuditLogService
}

func NewAuditLogHandler(service service.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{service: service}
}

func (h *AuditLogHandler) CreateAuditLog(c *fiber.Ctx) error {
	logrus.Info("Create audit log request received")

	var req request.CreateAuditLog
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create audit log request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateAuditLog(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create audit log")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("Audit log created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AuditLogHandler) GetAuditLogByID(c *fiber.Ctx) error {
	logrus.Info("Get audit log by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetAuditLogByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Audit log not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Audit log not found"})
	}

	logrus.WithField("auditLogID", idStr).Info("Audit log retrieved successfully")
	return c.JSON(resp)
}

func (h *AuditLogHandler) GetAuditLogsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get audit logs by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetAuditLogsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get audit logs")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Audit logs retrieved successfully")
	return c.JSON(resp)
}

func (h *AuditLogHandler) GetAllAuditLogs(c *fiber.Ctx) error {
	logrus.Info("Get all audit logs request received")

	resp, err := h.service.GetAllAuditLogs()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all audit logs")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.Info("All audit logs retrieved successfully")
	return c.JSON(resp)
}
