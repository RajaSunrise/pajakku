package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	CreateAttachment(attachment *models.Attachment) error
	GetAttachmentByID(id uint) (*models.Attachment, error)
	GetAttachmentsByUserID(userID uint) ([]*models.Attachment, error)
	UpdateAttachment(attachment *models.Attachment) error
	DeleteAttachment(id uint) error
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

func (r *attachmentRepository) CreateAttachment(attachment *models.Attachment) error {
	return r.db.Create(attachment).Error
}

func (r *attachmentRepository) GetAttachmentByID(id uint) (*models.Attachment, error) {
	var attachment models.Attachment
	err := r.db.First(&attachment, id).Error
	return &attachment, err
}

func (r *attachmentRepository) GetAttachmentsByUserID(userID uint) ([]*models.Attachment, error) {
	var attachments []*models.Attachment
	err := r.db.Where("user_id = ?", userID).Find(&attachments).Error
	return attachments, err
}

func (r *attachmentRepository) UpdateAttachment(attachment *models.Attachment) error {
	return r.db.Save(attachment).Error
}

func (r *attachmentRepository) DeleteAttachment(id uint) error {
	return r.db.Delete(&models.Attachment{}, id).Error
}
