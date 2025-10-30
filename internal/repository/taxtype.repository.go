package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type TaxTypeRepository interface {
	CreateTaxType(taxType *models.TaxType) error
	GetTaxTypeByID(id uint) (*models.TaxType, error)
	GetTaxTypeByCode(code string) (*models.TaxType, error)
	UpdateTaxType(taxType *models.TaxType) error
	DeleteTaxType(id uint) error
	GetAllTaxTypes() ([]models.TaxType, error)
}

type taxTypeRepository struct {
	db *gorm.DB
}

func NewTaxTypeRepository(db *gorm.DB) TaxTypeRepository {
	return &taxTypeRepository{db: db}
}

func (r *taxTypeRepository) CreateTaxType(taxType *models.TaxType) error {
	return r.db.Create(taxType).Error
}

func (r *taxTypeRepository) GetTaxTypeByID(id uint) (*models.TaxType, error) {
	var taxType models.TaxType
	err := r.db.First(&taxType, id).Error
	return &taxType, err
}

func (r *taxTypeRepository) GetTaxTypeByCode(code string) (*models.TaxType, error) {
	var taxType models.TaxType
	err := r.db.Where("kode_pajak = ?", code).First(&taxType).Error
	return &taxType, err
}

func (r *taxTypeRepository) UpdateTaxType(taxType *models.TaxType) error {
	return r.db.Save(taxType).Error
}

func (r *taxTypeRepository) DeleteTaxType(id uint) error {
	return r.db.Delete(&models.TaxType{}, id).Error
}

func (r *taxTypeRepository) GetAllTaxTypes() ([]models.TaxType, error) {
	var taxTypes []models.TaxType
	err := r.db.Find(&taxTypes).Error
	return taxTypes, err
}
