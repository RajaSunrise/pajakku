package repository

import (
	"github.com/RajaSunrise/pajakku/internal/models"

	"gorm.io/gorm"
)

type UsersAuthRepository interface {
	CreateUser(user *models.UserAuth) error
	GetUserByEmail(email string) (*models.UserAuth, error)
	GetUserByID(id string) (*models.UserAuth, error)
	UpdateUser(user *models.UserAuth) error
	UpdateUserProfileID(userID string, profileID uint) error
	DeleteUser(id string) error
	CreatePasswordResetToken(token *models.PasswordResetToken) error
	GetPasswordResetToken(token string) (*models.PasswordResetToken, error)
	DeletePasswordResetToken(token string) error
}

type usersAuthRepository struct {
	db *gorm.DB
}

func NewUsersAuthRepository(db *gorm.DB) UsersAuthRepository {
	return &usersAuthRepository{db: db}
}

func (r *usersAuthRepository) CreateUser(user *models.UserAuth) error {
	return r.db.Create(user).Error
}

func (r *usersAuthRepository) GetUserByEmail(email string) (*models.UserAuth, error) {
	var user models.UserAuth
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *usersAuthRepository) GetUserByID(id string) (*models.UserAuth, error) {
	var user models.UserAuth
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *usersAuthRepository) UpdateUser(user *models.UserAuth) error {
	return r.db.Save(user).Error
}

func (r *usersAuthRepository) UpdateUserProfileID(userID string, profileID uint) error {
	return r.db.Model(&models.UserAuth{}).Where("id = ?", userID).Update("user_profile_id", profileID).Error
}

func (r *usersAuthRepository) DeleteUser(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.UserAuth{}).Error
}

func (r *usersAuthRepository) CreatePasswordResetToken(token *models.PasswordResetToken) error {
	return r.db.Create(token).Error
}

func (r *usersAuthRepository) GetPasswordResetToken(token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken
	err := r.db.Where("token = ?", token).First(&resetToken).Error
	return &resetToken, err
}

func (r *usersAuthRepository) DeletePasswordResetToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.PasswordResetToken{}).Error
}
