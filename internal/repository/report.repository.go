package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type ReportRepository interface {
	CreateReport(report *models.ReportSPT) error
	GetReportByID(id uint) (*models.ReportSPT, error)
	GetReportsByUserID(userID string) ([]*models.ReportSPT, error)
	GetReportByNTTE(ntte string) (*models.ReportSPT, error)
	UpdateReport(report *models.ReportSPT) error
	DeleteReport(id uint) error
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) CreateReport(report *models.ReportSPT) error {
	return r.db.Create(report).Error
}

func (r *reportRepository) GetReportByID(id uint) (*models.ReportSPT, error) {
	var report models.ReportSPT
	err := r.db.First(&report, id).Error
	return &report, err
}

func (r *reportRepository) GetReportsByUserID(userID string) ([]*models.ReportSPT, error) {
	var reports []*models.ReportSPT
	err := r.db.Where("user_profile_id = ?", userID).Find(&reports).Error
	return reports, err
}

func (r *reportRepository) GetReportByNTTE(ntte string) (*models.ReportSPT, error) {
	var report models.ReportSPT
	err := r.db.Where("ntte = ?", ntte).First(&report).Error
	return &report, err
}

func (r *reportRepository) UpdateReport(report *models.ReportSPT) error {
	return r.db.Save(report).Error
}

func (r *reportRepository) DeleteReport(id uint) error {
	return r.db.Delete(&models.ReportSPT{}, id).Error
}
