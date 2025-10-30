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

type NotificationService interface {
	CreateNotification(userID string, req *request.CreateNotification) (*response.NotificationResponse, error)
	GetNotificationByID(id string) (*response.NotificationResponse, error)
	GetNotificationsByUserID(userID string) ([]*response.NotificationResponse, error)
	UpdateNotification(id string, req *request.UpdateNotification) (*response.NotificationResponse, error)
	DeleteNotification(id string) error
}

type notificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) CreateNotification(userID string, req *request.CreateNotification) (*response.NotificationResponse, error) {
	logrus.WithFields(logrus.Fields{"userID": userID, "judul": req.Judul}).Info("Create notification service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	notification := &models.Notification{
		UserID:       uint(userIDUint),
		Judul:        req.Judul,
		Isi:          req.Isi,
		Tipe:         req.Tipe,
		StatusBaca:   req.StatusBaca,
		TanggalKirim: req.TanggalKirim,
		TaxReturnID:  req.TaxReturnID,
		PaymentID:    req.PaymentID,
	}

	err = s.repo.CreateNotification(notification)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to create notification")
		return nil, err
	}

	logrus.WithField("userID", userID).Info("Notification created successfully")

	return &response.NotificationResponse{
		ID:           notification.ID,
		UserID:       notification.UserID,
		Judul:        notification.Judul,
		Isi:          notification.Isi,
		Tipe:         notification.Tipe,
		StatusBaca:   notification.StatusBaca,
		TanggalKirim: notification.TanggalKirim,
		TaxReturnID:  notification.TaxReturnID,
		PaymentID:    notification.PaymentID,
		CreatedAt:    notification.CreatedAt,
		UpdatedAt:    notification.UpdatedAt,
	}, nil
}

func (s *notificationService) GetNotificationByID(id string) (*response.NotificationResponse, error) {
	logrus.WithField("notificationID", id).Info("Get notification by ID service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid notification ID")
		return nil, errors.New("invalid notification ID")
	}

	notification, err := s.repo.GetNotificationByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("notificationID", id).Warn("Notification not found")
		return nil, err
	}

	logrus.WithField("notificationID", id).Info("Notification retrieved successfully")

	return &response.NotificationResponse{
		ID:           notification.ID,
		UserID:       notification.UserID,
		Judul:        notification.Judul,
		Isi:          notification.Isi,
		Tipe:         notification.Tipe,
		StatusBaca:   notification.StatusBaca,
		TanggalKirim: notification.TanggalKirim,
		TaxReturnID:  notification.TaxReturnID,
		PaymentID:    notification.PaymentID,
		CreatedAt:    notification.CreatedAt,
		UpdatedAt:    notification.UpdatedAt,
	}, nil
}

func (s *notificationService) GetNotificationsByUserID(userID string) ([]*response.NotificationResponse, error) {
	logrus.WithField("userID", userID).Info("Get notifications by user ID service called")

	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Invalid user ID")
		return nil, errors.New("invalid user ID")
	}

	notifications, err := s.repo.GetNotificationsByUserID(uint(userIDUint))
	if err != nil {
		logrus.WithError(err).WithField("userID", userID).Error("Failed to get notifications")
		return nil, err
	}

	var responses []*response.NotificationResponse
	for _, notification := range notifications {
		responses = append(responses, &response.NotificationResponse{
			ID:           notification.ID,
			UserID:       notification.UserID,
			Judul:        notification.Judul,
			Isi:          notification.Isi,
			Tipe:         notification.Tipe,
			StatusBaca:   notification.StatusBaca,
			TanggalKirim: notification.TanggalKirim,
			TaxReturnID:  notification.TaxReturnID,
			PaymentID:    notification.PaymentID,
			CreatedAt:    notification.CreatedAt,
			UpdatedAt:    notification.UpdatedAt,
		})
	}

	logrus.WithField("userID", userID).Info("Notifications retrieved successfully")

	return responses, nil
}

func (s *notificationService) UpdateNotification(id string, req *request.UpdateNotification) (*response.NotificationResponse, error) {
	logrus.WithField("notificationID", id).Info("Update notification service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid notification ID")
		return nil, errors.New("invalid notification ID")
	}

	notification, err := s.repo.GetNotificationByID(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("notificationID", id).Warn("Notification not found for update")
		return nil, err
	}

	if req.Judul != "" {
		notification.Judul = req.Judul
	}
	if req.Isi != "" {
		notification.Isi = req.Isi
	}
	if req.Tipe != "" {
		notification.Tipe = req.Tipe
	}
	if req.StatusBaca {
		notification.StatusBaca = req.StatusBaca
	}
	if !req.TanggalKirim.IsZero() {
		notification.TanggalKirim = req.TanggalKirim
	}
	if req.TaxReturnID != nil {
		notification.TaxReturnID = req.TaxReturnID
	}
	if req.PaymentID != nil {
		notification.PaymentID = req.PaymentID
	}

	err = s.repo.UpdateNotification(notification)
	if err != nil {
		logrus.WithError(err).WithField("notificationID", id).Error("Failed to update notification")
		return nil, err
	}

	logrus.WithField("notificationID", id).Info("Notification updated successfully")

	return &response.NotificationResponse{
		ID:           notification.ID,
		UserID:       notification.UserID,
		Judul:        notification.Judul,
		Isi:          notification.Isi,
		Tipe:         notification.Tipe,
		StatusBaca:   notification.StatusBaca,
		TanggalKirim: notification.TanggalKirim,
		TaxReturnID:  notification.TaxReturnID,
		PaymentID:    notification.PaymentID,
		CreatedAt:    notification.CreatedAt,
		UpdatedAt:    notification.UpdatedAt,
	}, nil
}

func (s *notificationService) DeleteNotification(id string) error {
	logrus.WithField("notificationID", id).Info("Delete notification service called")

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		logrus.WithError(err).WithField("id", id).Error("Invalid notification ID")
		return errors.New("invalid notification ID")
	}

	err = s.repo.DeleteNotification(uint(idUint))
	if err != nil {
		logrus.WithError(err).WithField("notificationID", id).Error("Failed to delete notification")
		return err
	}

	logrus.WithField("notificationID", id).Info("Notification deleted successfully")
	return nil
}
