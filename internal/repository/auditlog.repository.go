package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type AuditLogRepository interface {
	CreateAuditLog(auditLog *models.AuditLog) error
	GetAuditLogByID(id uint) (*models.AuditLog, error)
	GetAuditLogsByUserID(userID uint) ([]*models.AuditLog, error)
	GetAllAuditLogs() ([]models.AuditLog, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) CreateAuditLog(auditLog *models.AuditLog) error {
	return r.db.Create(auditLog).Error
}

func (r *auditLogRepository) GetAuditLogByID(id uint) (*models.AuditLog, error) {
	var auditLog models.AuditLog
	err := r.db.First(&auditLog, id).Error
	return &auditLog, err
}

func (r *auditLogRepository) GetAuditLogsByUserID(userID uint) ([]*models.AuditLog, error) {
	var auditLogs []*models.AuditLog
	err := r.db.Where("user_id = ?", userID).Find(&auditLogs).Error
	return auditLogs, err
}

func (r *auditLogRepository) GetAllAuditLogs() ([]models.AuditLog, error) {
	var auditLogs []models.AuditLog
	err := r.db.Find(&auditLogs).Error
	return auditLogs, err
}
