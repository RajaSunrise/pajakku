package service

import (
	"errors"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type RoleService interface {
	CreateRole(req *request.CreateRole) (*response.RoleResponse, error)
	GetRoleByID(id uint) (*response.RoleResponse, error)
	GetRoleByName(name string) (*response.RoleResponse, error)
	UpdateRole(id uint, req *request.UpdateRole) (*response.RoleResponse, error)
	DeleteRole(id uint) error
	GetAllRoles() ([]response.RoleResponse, error)
}

type roleService struct {
	repo repository.RoleRepository
}

func NewRoleService(repo repository.RoleRepository) RoleService {
	return &roleService{repo: repo}
}

func (s *roleService) CreateRole(req *request.CreateRole) (*response.RoleResponse, error) {
	logrus.WithField("namaRole", req.NamaRole).Info("Create role service called")

	// Check if role name already exists
	_, err := s.repo.GetRoleByName(req.NamaRole)
	if err == nil {
		logrus.WithField("namaRole", req.NamaRole).Warn("Role name already exists")
		return nil, errors.New("role name already exists")
	}

	role := &models.Role{
		NamaRole:    req.NamaRole,
		Permissions: req.Permissions,
	}

	err = s.repo.CreateRole(role)
	if err != nil {
		logrus.WithError(err).Error("Failed to create role")
		return nil, err
	}

	logrus.WithField("roleID", role.ID).Info("Role created successfully")

	return &response.RoleResponse{
		ID:          role.ID,
		NamaRole:    role.NamaRole,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) GetRoleByID(id uint) (*response.RoleResponse, error) {
	logrus.WithField("id", id).Info("Get role by ID service called")

	role, err := s.repo.GetRoleByID(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Warn("Role not found")
		return nil, err
	}

	logrus.WithField("id", id).Info("Role retrieved successfully")

	return &response.RoleResponse{
		ID:          role.ID,
		NamaRole:    role.NamaRole,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) GetRoleByName(name string) (*response.RoleResponse, error) {
	logrus.WithField("name", name).Info("Get role by name service called")

	role, err := s.repo.GetRoleByName(name)
	if err != nil {
		logrus.WithError(err).WithField("name", name).Warn("Role not found")
		return nil, err
	}

	logrus.WithField("name", name).Info("Role retrieved successfully")

	return &response.RoleResponse{
		ID:          role.ID,
		NamaRole:    role.NamaRole,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) UpdateRole(id uint, req *request.UpdateRole) (*response.RoleResponse, error) {
	logrus.WithField("id", id).Info("Update role service called")

	role, err := s.repo.GetRoleByID(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Warn("Role not found for update")
		return nil, err
	}

	if req.NamaRole != "" {
		// Check if new name already exists and not the current role
		existingRole, err := s.repo.GetRoleByName(req.NamaRole)
		if err == nil && existingRole.ID != id {
			logrus.WithField("namaRole", req.NamaRole).Warn("Role name already exists")
			return nil, errors.New("role name already exists")
		}
		role.NamaRole = req.NamaRole
	}

	if req.Permissions != "" {
		role.Permissions = req.Permissions
	}

	err = s.repo.UpdateRole(role)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Failed to update role")
		return nil, err
	}

	logrus.WithField("id", id).Info("Role updated successfully")

	return &response.RoleResponse{
		ID:          role.ID,
		NamaRole:    role.NamaRole,
		Permissions: role.Permissions,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (s *roleService) DeleteRole(id uint) error {
	logrus.WithField("id", id).Info("Delete role service called")

	err := s.repo.DeleteRole(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Failed to delete role")
		return err
	}

	logrus.WithField("id", id).Info("Role deleted successfully")
	return nil
}

func (s *roleService) GetAllRoles() ([]response.RoleResponse, error) {
	logrus.Info("Get all roles service called")

	roles, err := s.repo.GetAllRoles()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all roles")
		return nil, err
	}

	var roleResponses []response.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, response.RoleResponse{
			ID:          role.ID,
			NamaRole:    role.NamaRole,
			Permissions: role.Permissions,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}

	logrus.Info("All roles retrieved successfully")
	return roleResponses, nil
}
