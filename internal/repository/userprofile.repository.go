package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	CreateProfile(profile *models.UserProfile) error
	GetProfileByID(id uint) (*models.UserProfile, error)
	GetProfileByNPWP(npwp string) (*models.UserProfile, error)
	UpdateProfile(profile *models.UserProfile) error
	DeleteProfile(id uint) error
}

type userProfileRepository struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepository{db: db}
}

func (r *userProfileRepository) CreateProfile(profile *models.UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *userProfileRepository) GetProfileByID(id uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.First(&profile, id).Error
	return &profile, err
}

func (r *userProfileRepository) GetProfileByNPWP(npwp string) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Where("npwp = ?", npwp).First(&profile).Error
	return &profile, err
}

func (r *userProfileRepository) UpdateProfile(profile *models.UserProfile) error {
	return r.db.Save(profile).Error
}

func (r *userProfileRepository) DeleteProfile(id uint) error {
	return r.db.Delete(&models.UserProfile{}, id).Error
}
