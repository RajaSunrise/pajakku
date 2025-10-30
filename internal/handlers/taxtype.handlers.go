package handlers

import (
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TaxTypeHandler struct {
	service service.TaxTypeService
}

func NewTaxTypeHandler(service service.TaxTypeService) *TaxTypeHandler {
	return &TaxTypeHandler{service: service}
}

func (h *TaxTypeHandler) CreateTaxType(c *fiber.Ctx) error {
	logrus.Info("Create tax type request received")

	var req request.CreateTaxType
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create tax type request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateTaxType(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create tax type")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("Tax type created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *TaxTypeHandler) GetTaxTypeByID(c *fiber.Ctx) error {
	logrus.Info("Get tax type by ID request received")
	idStr := c.Params("id")

	resp, err := h.service.GetTaxTypeByID(idStr)
	if err != nil {
		logrus.WithError(err).Warn("Tax type not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tax type not found"})
	}

	logrus.WithField("taxTypeID", idStr).Info("Tax type retrieved successfully")
	return c.JSON(resp)
}

func (h *TaxTypeHandler) GetTaxTypeByCode(c *fiber.Ctx) error {
	logrus.Info("Get tax type by code request received")
	code := c.Params("code")

	resp, err := h.service.GetTaxTypeByCode(code)
	if err != nil {
		logrus.WithError(err).Warn("Tax type not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tax type not found"})
	}

	logrus.WithField("code", code).Info("Tax type retrieved successfully")
	return c.JSON(resp)
}

func (h *TaxTypeHandler) UpdateTaxType(c *fiber.Ctx) error {
	logrus.Info("Update tax type request received")
	idStr := c.Params("id")

	var req request.UpdateTaxType
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update tax type request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateTaxType(idStr, &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update tax type")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("taxTypeID", idStr).Info("Tax type updated successfully")
	return c.JSON(resp)
}

func (h *TaxTypeHandler) DeleteTaxType(c *fiber.Ctx) error {
	logrus.Info("Delete tax type request received")
	idStr := c.Params("id")

	err := h.service.DeleteTaxType(idStr)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete tax type")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("taxTypeID", idStr).Info("Tax type deleted successfully")
	return c.JSON(fiber.Map{"message": "Tax type deleted successfully"})
}

func (h *TaxTypeHandler) GetAllTaxTypes(c *fiber.Ctx) error {
	logrus.Info("Get all tax types request received")

	resp, err := h.service.GetAllTaxTypes()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all tax types")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.Info("All tax types retrieved successfully")
	return c.JSON(resp)
}
