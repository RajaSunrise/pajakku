package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	GetPaymentByID(id uint) (*models.Payment, error)
	GetPaymentsByUserID(userID uint) ([]*models.Payment, error)
	UpdatePayment(payment *models.Payment) error
	DeletePayment(id uint) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentByID(id uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.First(&payment, id).Error
	return &payment, err
}

func (r *paymentRepository) GetPaymentsByUserID(userID uint) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.Where("user_id = ?", userID).Find(&payments).Error
	return payments, err
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) DeletePayment(id uint) error {
	return r.db.Delete(&models.Payment{}, id).Error
}
