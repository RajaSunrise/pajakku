package service

import (
	"errors"
	"time"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/RajaSunrise/pajakku/pkg/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersAuthService interface {
	Register(req *request.CreateUsersAuth) error
	Login(req *request.LoginRequest) (*response.LoginResponse, error)
	ForgetPassword(req *request.ForgetPasswordRequest) error
	ResetPassword(req *request.ResetPasswordRequest) error
}

type usersAuthService struct {
	repo repository.UsersAuthRepository
}

func NewUsersAuthService(repo repository.UsersAuthRepository) UsersAuthService {
	return &usersAuthService{repo: repo}
}

func (s *usersAuthService) Register(req *request.CreateUsersAuth) error {
	// Check if user already exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.UserAuth{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	return s.repo.CreateUser(user)
}

func (s *usersAuthService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	now := time.Now()
	user.TerakhirLogin = &now
	s.repo.UpdateUser(user)

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &response.LoginResponse{
		Token:     token,
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}

func (s *usersAuthService) ForgetPassword(req *request.ForgetPasswordRequest) error {
	// Check if user exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// Generate token
	token := uuid.New().String()
	expiresAt := time.Now().Add(time.Hour * 1)

	resetToken := &models.PasswordResetToken{
		Email:     req.Email,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	return s.repo.CreatePasswordResetToken(resetToken)
}

func (s *usersAuthService) ResetPassword(req *request.ResetPasswordRequest) error {
	resetToken, err := s.repo.GetPasswordResetToken(req.Token)
	if err != nil {
		return errors.New("invalid token")
	}

	if time.Now().After(resetToken.ExpiresAt) {
		return errors.New("token expired")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := s.repo.GetUserByEmail(resetToken.Email)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = s.repo.UpdateUser(user)
	if err != nil {
		return err
	}

	// Delete token
	return s.repo.DeletePasswordResetToken(req.Token)
}
