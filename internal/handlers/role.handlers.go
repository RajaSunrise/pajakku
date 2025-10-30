package handlers

import (
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RoleHandler struct {
	service service.RoleService
}

func NewRoleHandler(service service.RoleService) *RoleHandler {
	return &RoleHandler{service: service}
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	logrus.Info("Create role request received")

	var req request.CreateRole
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse create role request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.CreateRole(&req)
	if err != nil {
		logrus.WithError(err).Error("Failed to create role")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.Info("Role created successfully")
	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *RoleHandler) GetRoleByID(c *fiber.Ctx) error {
	logrus.Info("Get role by ID request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid role ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	resp, err := h.service.GetRoleByID(uint(id))
	if err != nil {
		logrus.WithError(err).Warn("Role not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	logrus.WithField("roleID", id).Info("Role retrieved successfully")
	return c.JSON(resp)
}

func (h *RoleHandler) GetRoleByName(c *fiber.Ctx) error {
	logrus.Info("Get role by name request received")
	name := c.Params("name")

	resp, err := h.service.GetRoleByName(name)
	if err != nil {
		logrus.WithError(err).Warn("Role not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Role not found"})
	}

	logrus.WithField("name", name).Info("Role retrieved successfully")
	return c.JSON(resp)
}

func (h *RoleHandler) GetAllRoles(c *fiber.Ctx) error {
	logrus.Info("Get all roles request received")

	resp, err := h.service.GetAllRoles()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all roles")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.Info("All roles retrieved successfully")
	return c.JSON(resp)
}

func (h *RoleHandler) UpdateRole(c *fiber.Ctx) error {
	logrus.Info("Update role request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid role ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	var req request.UpdateRole
	if err := c.BodyParser(&req); err != nil {
		logrus.WithError(err).Warn("Failed to parse update role request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	resp, err := h.service.UpdateRole(uint(id), &req)
	if err != nil {
		logrus.WithError(err).Error("Failed to update role")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	logrus.WithField("roleID", id).Info("Role updated successfully")
	return c.JSON(resp)
}

func (h *RoleHandler) DeleteRole(c *fiber.Ctx) error {
	logrus.Info("Delete role request received")
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		logrus.WithError(err).Warn("Invalid role ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid role ID"})
	}

	err = h.service.DeleteRole(uint(id))
	if err != nil {
		logrus.WithError(err).Error("Failed to delete role")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
	}

	logrus.WithField("roleID", id).Info("Role deleted successfully")
	return c.JSON(fiber.Map{"message": "Role deleted successfully"})
}
