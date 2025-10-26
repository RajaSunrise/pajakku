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
	"github.com/sirupsen/logrus"
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
	logrus.WithField("email", req.Email).Info("Register service called")

	// Check if user already exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err == nil {
		logrus.WithField("email", req.Email).Warn("User already exists")
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to hash password")
		return err
	}

	user := &models.UserAuth{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to create user")
		return err
	}

	logrus.WithField("email", req.Email).Info("User registered successfully")
	return nil
}

func (s *usersAuthService) Login(req *request.LoginRequest) (*response.LoginResponse, error) {
	logrus.WithField("email", req.Email).Info("Login service called")

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		logrus.WithField("email", req.Email).Warn("User not found or invalid credentials")
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logrus.WithField("email", req.Email).Warn("Invalid password")
		return nil, errors.New("invalid credentials")
	}

	// Update last login
	now := time.Now()
	user.TerakhirLogin = &now
	err = s.repo.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to update last login")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to generate token")
		return nil, err
	}

	logrus.WithField("email", req.Email).Info("User logged in successfully")
	return &response.LoginResponse{
		Token:     token,
		UserID:    user.ID,
		Email:     user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}, nil
}

func (s *usersAuthService) ForgetPassword(req *request.ForgetPasswordRequest) error {
	logrus.WithField("email", req.Email).Info("Forget password service called")

	// Check if user exists
	_, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		logrus.WithField("email", req.Email).Warn("User not found for password reset")
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

	err = s.repo.CreatePasswordResetToken(resetToken)
	if err != nil {
		logrus.WithError(err).WithField("email", req.Email).Error("Failed to create password reset token")
		return err
	}

	logrus.WithField("email", req.Email).Info("Password reset token created")
	return nil
}

func (s *usersAuthService) ResetPassword(req *request.ResetPasswordRequest) error {
	logrus.Info("Reset password service called")

	resetToken, err := s.repo.GetPasswordResetToken(req.Token)
	if err != nil {
		logrus.WithError(err).Warn("Invalid reset token")
		return errors.New("invalid token")
	}

	if time.Now().After(resetToken.ExpiresAt) {
		logrus.WithField("email", resetToken.Email).Warn("Reset token expired")
		return errors.New("token expired")
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).WithField("email", resetToken.Email).Error("Failed to hash new password")
		return err
	}

	user, err := s.repo.GetUserByEmail(resetToken.Email)
	if err != nil {
		logrus.WithError(err).WithField("email", resetToken.Email).Error("Failed to get user for password reset")
		return err
	}

	user.Password = string(hashedPassword)
	err = s.repo.UpdateUser(user)
	if err != nil {
		logrus.WithError(err).WithField("email", resetToken.Email).Error("Failed to update user password")
		return err
	}

	// Delete token
	err = s.repo.DeletePasswordResetToken(req.Token)
	if err != nil {
		logrus.WithError(err).WithField("email", resetToken.Email).Error("Failed to delete reset token")
		return err
	}

	logrus.WithField("email", resetToken.Email).Info("Password reset successfully")
	return nil
}
