package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type BillingRepository interface {
	CreateBilling(billing *models.Billing) error
	GetBillingByID(id string) (*models.Billing, error)
	GetBillingsByUserID(userID string) ([]*models.Billing, error)
	GetBillingByKodeBilling(kodeBilling string) (*models.Billing, error)
	UpdateBilling(billing *models.Billing) error
	DeleteBilling(id string) error
}

type billingRepository struct {
	db *gorm.DB
}

func NewBillingRepository(db *gorm.DB) BillingRepository {
	return &billingRepository{db: db}
}

func (r *billingRepository) CreateBilling(billing *models.Billing) error {
	return r.db.Create(billing).Error
}

func (r *billingRepository) GetBillingByID(id string) (*models.Billing, error) {
	var billing models.Billing
	err := r.db.First(&billing, id).Error
	return &billing, err
}

func (r *billingRepository) GetBillingsByUserID(userID string) ([]*models.Billing, error) {
	var billings []*models.Billing
	err := r.db.Where("user_profile_id = ?", userID).Find(&billings).Error
	return billings, err
}

func (r *billingRepository) GetBillingByKodeBilling(kodeBilling string) (*models.Billing, error) {
	var billing models.Billing
	err := r.db.Where("kode_billing = ?", kodeBilling).First(&billing).Error
	return &billing, err
}

func (r *billingRepository) UpdateBilling(billing *models.Billing) error {
	return r.db.Save(billing).Error
}

func (r *billingRepository) DeleteBilling(id string) error {
	return r.db.Delete(&models.Billing{}, id).Error
}
