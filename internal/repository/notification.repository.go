package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(notification *models.Notification) error
	GetNotificationByID(id uint) (*models.Notification, error)
	GetNotificationsByUserID(userID uint) ([]*models.Notification, error)
	UpdateNotification(notification *models.Notification) error
	DeleteNotification(id uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateNotification(notification *models.Notification) error {
	return r.db.Create(notification).Error
}

func (r *notificationRepository) GetNotificationByID(id uint) (*models.Notification, error) {
	var notification models.Notification
	err := r.db.First(&notification, id).Error
	return &notification, err
}

func (r *notificationRepository) GetNotificationsByUserID(userID uint) ([]*models.Notification, error) {
	var notifications []*models.Notification
	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) UpdateNotification(notification *models.Notification) error {
	return r.db.Save(notification).Error
}

func (r *notificationRepository) DeleteNotification(id uint) error {
	return r.db.Delete(&models.Notification{}, id).Error
}
