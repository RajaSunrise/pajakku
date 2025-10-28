package handlers

import (
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ReportHandler struct {
	service service.ReportService
}

func NewReportHandler(service service.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) CreateReport(c *fiber.Ctx) error {
	logrus.Info("Create report request received")
	userID := c.Locals("userID").(string)

	var req request.ReportSPTRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create report request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateReport(userID, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create report")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("userID", userID).Info("Report created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *ReportHandler) GetReportByID(c *fiber.Ctx) error {
	logrus.Info("Get report by ID request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid report ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	resp, err := h.service.GetReportByID(uint(id))
	if err != nil {
		logrus.WithError(err).Warn("Report not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}

	logrus.WithField("reportID", id).Info("Report retrieved successfully")
	return c.JSON(resp)
}

func (h *ReportHandler) GetReportsByUserID(c *fiber.Ctx) error {
	logrus.Info("Get reports by user ID request received")
	userID := c.Locals("userID").(string)

	resp, err := h.service.GetReportsByUserID(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to get reports")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("userID", userID).Info("Reports retrieved successfully")
	return c.JSON(resp)
}

func (h *ReportHandler) UpdateReport(c *fiber.Ctx) error {
	logrus.Info("Update report request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid report ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	var req request.ReportSPTRequest
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update report request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateReport(uint(id), &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update report")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("reportID", id).Info("Report updated successfully")
	return c.JSON(resp)
}

func (h *ReportHandler) DeleteReport(c *fiber.Ctx) error {
	logrus.Info("Delete report request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid report ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid report ID"})
	}

	err = h.service.DeleteReport(uint(id))
	if err != nil {
		logrus.WithError(err).Error("Failed to delete report")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("reportID", id).Info("Report deleted successfully")
	return c.JSON(fiber.Map{"message": "Report deleted successfully"})
}
