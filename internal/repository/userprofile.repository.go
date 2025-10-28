package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"
	"gorm.io/gorm"
)

type UserProfileRepository interface {
	CreateProfile(profile *models.UserProfile) error
	GetProfileByNIK(nik uint) (*models.UserProfile, error)
	GetProfileByUserID(userID string) (*models.UserProfile, error)
	GetProfileByNPWP(npwp string) (*models.UserProfile, error)
	UpdateProfile(profile *models.UserProfile) error
	DeleteProfileByNIK(nik uint) error
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

func (r *userProfileRepository) GetProfileByNIK(nik uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.First(&profile, nik).Error
	return &profile, err
}

func (r *userProfileRepository) GetProfileByUserID(userID string) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Joins("UserAuth").Where(`"UserAuth".id = ?`, userID).First(&profile).Error
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

func (r *userProfileRepository) DeleteProfileByNIK(nik uint) error {
	return r.db.Delete(&models.UserProfile{}, nik).Error
}
