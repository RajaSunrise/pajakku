package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type InvoiceRepository interface {
	CreateInvoice(invoice *models.Invoice) error
	GetInvoiceByID(id uint) (*models.Invoice, error)
	GetInvoicesByUserID(userID uint) ([]*models.Invoice, error)
	UpdateInvoice(invoice *models.Invoice) error
	DeleteInvoice(id uint) error
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) CreateInvoice(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *invoiceRepository) GetInvoiceByID(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.First(&invoice, id).Error
	return &invoice, err
}

func (r *invoiceRepository) GetInvoicesByUserID(userID uint) ([]*models.Invoice, error) {
	var invoices []*models.Invoice
	err := r.db.Where("user_id = ?", userID).Find(&invoices).Error
	return invoices, err
}

func (r *invoiceRepository) UpdateInvoice(invoice *models.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *invoiceRepository) DeleteInvoice(id uint) error {
	return r.db.Delete(&models.Invoice{}, id).Error
}
