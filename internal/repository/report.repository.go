package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type TaxReturnRepository interface {
	CreateTaxReturn(taxReturn *models.TaxReturn) error
	GetTaxReturnByID(id uint) (*models.TaxReturn, error)
	GetTaxReturnsByUserID(userID uint) ([]*models.TaxReturn, error)
	UpdateTaxReturn(taxReturn *models.TaxReturn) error
	DeleteTaxReturn(id uint) error
}

type taxReturnRepository struct {
	db *gorm.DB
}

func NewTaxReturnRepository(db *gorm.DB) TaxReturnRepository {
	return &taxReturnRepository{db: db}
}

func (r *taxReturnRepository) CreateTaxReturn(taxReturn *models.TaxReturn) error {
	return r.db.Create(taxReturn).Error
}

func (r *taxReturnRepository) GetTaxReturnByID(id uint) (*models.TaxReturn, error) {
	var taxReturn models.TaxReturn
	err := r.db.First(&taxReturn, id).Error
	return &taxReturn, err
}

func (r *taxReturnRepository) GetTaxReturnsByUserID(userID uint) ([]*models.TaxReturn, error) {
	var taxReturns []*models.TaxReturn
	err := r.db.Where("user_id = ?", userID).Find(&taxReturns).Error
	return taxReturns, err
}

func (r *taxReturnRepository) UpdateTaxReturn(taxReturn *models.TaxReturn) error {
	return r.db.Save(taxReturn).Error
}

func (r *taxReturnRepository) DeleteTaxReturn(id uint) error {
	return r.db.Delete(&models.TaxReturn{}, id).Error
}
