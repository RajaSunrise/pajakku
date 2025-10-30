package service

import (
	"errors"
	"strconv"

	"github.com/RajaSunrise/pajakku/internal/models"
	"github.com/RajaSunrise/pajakku/internal/models/request"
	"github.com/RajaSunrise/pajakku/internal/models/response"
	"github.com/RajaSunrise/pajakku/internal/repository"
	"github.com/sirupsen/logrus"
)

type AuditLogService interface {
	CreateAuditLog(req *request.CreateAuditLog) (*response.AuditLogResponse, error)
	GetAuditLogByID(id string) (*response.AuditLogResponse, error)
	GetAuditLogsByUserID(userID string) ([]*response.AuditLogResponse, error)
	GetAllAuditLogs() ([]*response.AuditLogResponse, error)
}

type auditLogService struct {
	repo repository.AuditLogRepository
}

func NewAuditLogService(repo repository.AuditLogRepository) AuditLogService {
	return &auditLogService{repo: repo}
}

func (s *auditLogService) CreateAuditLog(req *request.CreateAuditLog) (*response.AuditLogResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": req.UserID, "aksi": req.Aksi}).Info("Create audit log service called")

	auditLog := &models.AuditLog{
		UserID:    req.UserID,
		Aksi:      req.Aksi,
		IPAddress: req.IPAddress,
		OldValue:  req.OldValue,
		NewValue:  req.NewValue,
	}

	err := s.repo.CreateAuditLog(auditLog)
	if err != nil {
		logrus.WithError(err).WithField("userID", req.UserID).Error("Failed to create audit log")
		return nil, err
	}

	logrus.WithField("userID", req.UserID).Info("Audit log created successfully")

	return &response.AuditLogResponse{
		ID:        auditLog.ID,
		UserID:    auditLog.UserID,
		Aksi:      auditLog.Aksi,
		Timestamp: auditLog.Timestamp,
		IPAddress: auditLog.IPAddress,
		OldValue:  auditLog.OldValue,
		NewValue:  auditLog.NewValue,
		CreatedAt: auditLog.CreatedAt,
		UpdatedAt: auditLog.UpdatedAt,
	}, nil
}

func (s *auditLogService) GetAuditLogByID(id string) (*response.AuditLogResponse, error) {
	logrus.WithField("auditLogID", id).Info("Get audit log by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid audit log ID")
		return nil, errors.New("invalid audit log ID")
	}

	auditLog, err := s.repo.GetAuditLogByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("auditLogID", id).Warn("Audit log not found")
		return nil, err
	}

	logrus.WithField("auditLogID", id).Info("Audit log retrieved successfully")

	return &response.AuditLogResponse{
		ID:        auditLog.ID,
		UserID:    auditLog.UserID,
		Aksi:      auditLog.Aksi,
		Timestamp: auditLog.Timestamp,
		IPAddress: auditLog.IPAddress,
		OldValue:  auditLog.OldValue,
		NewValue:  auditLog.NewValue,
		CreatedAt: auditLog.CreatedAt,
		UpdatedAt: auditLog.UpdatedAt,
	}, nil
}

func (s *auditLogService) GetAuditLogsByUserID(userID string) ([]*response.AuditLogResponse, error) {
	logrus.WithField("userID", userID).Info("Get audit logs by user ID service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	auditLogs, err := s.repo.GetAuditLogsByUserID(uint(userIDUint))
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get audit logs")
		return nil, err
	}

	var responses []*response.AuditLogResponse
	for _, auditLog := range auditLogs {
		responses = append(responses, &response.AuditLogResponse{
			ID:        auditLog.ID,
			UserID:    auditLog.UserID,
			Aksi:      auditLog.Aksi,
			Timestamp: auditLog.Timestamp,
			IPAddress: auditLog.IPAddress,
			OldValue:  auditLog.OldValue,
			NewValue:  auditLog.NewValue,
			CreatedAt: auditLog.CreatedAt,
			UpdatedAt: auditLog.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Audit logs retrieved successfully")

	return responses, nil
}

func (s *auditLogService) GetAllAuditLogs() ([]*response.AuditLogResponse, error) {
	logrus.Info("Get all audit logs service called")

	auditLogs, err := s.repo.GetAllAuditLogs()
	if err != nil {
		logrus.WithError(err).Error("Failed to get all audit logs")
		return nil, err
	}

	var responses []*response.AuditLogResponse
	for _, auditLog := range auditLogs {
		responses = append(responses, &response.AuditLogResponse{
			ID:        auditLog.ID,
			UserID:    auditLog.UserID,
			Aksi:      auditLog.Aksi,
			Timestamp: auditLog.Timestamp,
			IPAddress: auditLog.IPAddress,
			OldValue:  auditLog.OldValue,
			NewValue:  auditLog.NewValue,
			CreatedAt: auditLog.CreatedAt,
			UpdatedAt: auditLog.UpdatedAt,
		})
	}

	logrus.Info("All audit logs retrieved successfully")

	return responses, nil
}
