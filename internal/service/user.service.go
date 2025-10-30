package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/RajaSunrise/pajakku/internal/databases"
	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/pkg/utils"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(req *request.CreateUser) (*response.UserResponse, error)
	GetUserByID(id uint) (*response.UserResponse, error)
	GetUserByEmail(email string) (*response.UserResponse, error)
	GetAllUsers() ([]response.UserResponse, error)
	UpdateUser(id uint, req *request.UpdateUser) (*response.UserResponse, error)
	DeleteUser(id uint) error
	Login(req *request.LoginRequest) (*response.LoginResponse, error)
}

type userService struct {
	repo     repository.UserRepository
	roleRepo repository.RoleRepository
}

func NewUserService(repo repository.UserRepository, roleRepo repository.RoleRepository) UserService {
	return &userService{repo: repo, roleRepo: roleRepo}
}

func (s *userService) CreateUser(req *request.CreateUser) (*response.UserResponse, error) {
	logrus.WithFields(logrus.Fields{"email": req.Email, "npwp": req.NPWP}).Info("Create user service called")

	// Check if email already exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		logrus.WithField("email", req.Email).Warn("Email already exists")
		return nil, errors.New("email already exists")
	}

	// Check if NPWP already exists
	_, err = s.repo.GetUserByNPWP(req.NPWP)
	if err == nil {
		logrus.WithField("npwp", req.NPWP).Warn("NPWP already exists")
		return nil, errors.New("NPWP already exists")
	}

	// Check if role exists
	_, err = s.roleRepo.GetRoleByID(req.RoleID)
	if err != nil {
		logrus.WithError(err).WithField("roleID", req.RoleID).Error("Role not found")
		return nil, errors.New("role not found")
	}

	// Generate unique random ID
	var id uint
	for {
		id = utils.GenerateRandomID()
		_, err = s.repo.GetUserByID(id)
		if err != nil { // ID not found, so it's unique
			break
		}
	}

	user := &models.User{
		ID:              id,
		NIK:             req.NIK,
		NPWP:            req.NPWP,
		Nama:            req.Nama,
		Email:           req.Email,
		Alamat:          req.Alamat,
		JenisWajibPajak: req.JenisWajibPajak,
		RoleID:          req.RoleID,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		return nil, err
	}

	// Create UserAuth
	userAuth := &models.UserAuth{
		UserID: user.ID,
	}
	if err := userAuth.HashPassword(req.Password); err != nil {
		logrus.WithError(err).Error("Failed to hash password")
		return nil, err
	}
	if err := databases.DB.Create(userAuth).Error; err != nil {
		logrus.WithError(err).Error("Failed to create user auth")
		return nil, err
	}

	logrus.WithField("userID", user.ID).Info("User created successfully")

	return &response.UserResponse{
		ID:                user.ID,
		NIK:               user.NIK,
		NPWP:              user.NPWP,
		Nama:              user.Nama,
		Email:             user.Email,
		Alamat:            user.Alamat,
		JenisWajibPajak:   user.JenisWajibPajak,
		TanggalRegistrasi: user.TanggalRegistrasi,
		StatusAktif:       user.StatusAktif,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}, nil
}

func (s *userService) GetUserByID(id uint) (*response.UserResponse, error) {
	logrus.WithField("id", id).Info("Get user by ID service called")

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Warn("User not found")
		return nil, err
	}

	logrus.WithField("id", id).Info("User retrieved successfully")

	return &response.UserResponse{
		ID:                user.ID,
		NIK:               user.NIK,
		NPWP:              user.NPWP,
		Nama:              user.Nama,
		Email:             user.Email,
		Alamat:            user.Alamat,
		JenisWajibPajak:   user.JenisWajibPajak,
		TanggalRegistrasi: user.TanggalRegistrasi,
		StatusAktif:       user.StatusAktif,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetUserByEmail(email string) (*response.UserResponse, error) {
	logrus.WithField("email", email).Info("Get user by email service called")

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		logrus.WithError(err).WithField("email", email).Warn("User not found")
		return nil, err
	}

	logrus.WithField("email", email).Info("User retrieved successfully")

	return &response.UserResponse{
		ID:                user.ID,
		NIK:               user.NIK,
		NPWP:              user.NPWP,
		Nama:              user.Nama,
		Email:             user.Email,
		Alamat:            user.Alamat,
		JenisWajibPajak:   user.JenisWajibPajak,
		TanggalRegistrasi: user.TanggalRegistrasi,
		StatusAktif:       user.StatusAktif,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) UpdateUser(id uint, req *request.UpdateUser) (*response.UserResponse, error) {
	logrus.WithField("id", id).Info("Update user service called")

	user, err := s.repo.GetUserByID(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Warn("User not found for update")
		return nil, err
	}

	if req.NIK != 0 {
		user.NIK = req.NIK
	}
	if req.NPWP != "" {
		// Check if new NPWP already exists and not the current user
		existingUser, err := s.repo.GetUserByNPWP(req.NPWP)
		if err == nil && existingUser.ID != id {
			logrus.WithField("npwp", req.NPWP).Warn("NPWP already exists")
			return nil, errors.New("NPWP already exists")
		}
		user.NPWP = req.NPWP
	}

	if req.Email != "" {
		// Check if new email already exists and not the current user
		existingUser, err := s.repo.GetUserByEmail(req.Email)
		if err == nil && existingUser.ID != id {
			logrus.WithField("email", req.Email).Warn("Email already exists")
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	if req.Password != "" {
		// Update password in UserAuth
		var userAuth models.UserAuth
		if err := databases.DB.Where("user_id = ?", user.ID).First(&userAuth).Error; err != nil {
			logrus.WithError(err).Error("UserAuth not found")
			return nil, errors.New("user auth not found")
		}
		if err := userAuth.HashPassword(req.Password); err != nil {
			logrus.WithError(err).Error("Failed to hash password")
			return nil, err
		}
		if err := databases.DB.Save(&userAuth).Error; err != nil {
			logrus.WithError(err).Error("Failed to update user auth")
			return nil, err
		}
	}

	if req.Nama != "" {
		user.Nama = req.Nama
	}
	if req.Alamat != "" {
		user.Alamat = req.Alamat
	}
	if req.JenisWajibPajak != "" {
		user.JenisWajibPajak = req.JenisWajibPajak
	}
	if req.RoleID != 0 {
		// Check if role exists
		_, err := s.roleRepo.GetRoleByID(req.RoleID)
		if err != nil {
			logrus.WithError(err).WithField("roleID", req.RoleID).Error("Role not found")
			return nil, errors.New("role not found")
		}
		user.RoleID = req.RoleID
	}
	if req.StatusAktif != nil {
		user.StatusAktif = *req.StatusAktif
	}

	err = s.repo.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Failed to update user")
		return nil, err
	}

	logrus.WithField("id", id).Info("User updated successfully")

	return &response.UserResponse{
		ID:                user.ID,
		NIK:               user.NIK,
		NPWP:              user.NPWP,
		Nama:              user.Nama,
		Email:             user.Email,
		Alamat:            user.Alamat,
		JenisWajibPajak:   user.JenisWajibPajak,
		TanggalRegistrasi: user.TanggalRegistrasi,
		StatusAktif:       user.StatusAktif,

		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) DeleteUser(id uint) error {
	logrus.WithField("id", id).Info("Delete user service called")

	err := s.repo.DeleteUser(id)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Failed to delete user")
		return err
	}

	logrus.WithField("id", id).Info("User deleted successfully")
	return nil
}

func (s *userService) GetAllUsers() ([]response.UserResponse, error) {
	logrus.Info("Get all users service called")

	users, err := s.repo.GetAllUsers()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all users")
		return nil, err
	}

	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, response.UserResponse{
			ID:                user.ID,
			NIK:               user.NIK,
			NPWP:              user.NPWP,
			Nama:              user.Nama,
			Email:             user.Email,
			Alamat:            user.Alamat,
			JenisWajibPajak:   user.JenisWajibPajak,
			TanggalRegistrasi: user.TanggalRegistrasi,
			StatusAktif:       user.StatusAktif,

			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	logrus.Info("All users retrieved successfully")
	return userResponses, nil
}

func (s *userService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	logrus.WithField("email", req.Email).Info("Login service called")

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Warn("User not found")
		return nil, errors.New("invalid credentials")
	}

	// Check password
	var userAuth models.UserAuth
	if err := databases.DB.Where("user_id = ?", user.ID).First(&userAuth).Error; err != nil {
		logrus.WithError(err).Error("UserAuth not found")
		return nil, errors.New("invalid credentials")
	}

	if !userAuth.CheckPassword(req.Password) {
		logrus.WithField("email", req.Email).Warn("Invalid password")
		return nil, errors.New("invalid credentials")
	}

	// Generate token
	token, err := utils.GenerateToken(fmt.Sprintf("%d", user.ID), user.Email, user.RoleID)
	if err != nil {
		logrus.WithError(err).Error("Failed to generate token")
		return nil, err
	}

	logrus.WithField("email", req.Email).Info("Login successful")

	return &response.LoginResponse{
		Token:     token,
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}, nil
}
