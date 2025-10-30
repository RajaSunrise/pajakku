package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserAuth model - Model untuk autentikasi pengguna.
type UserAuth struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UserID    uint   `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relasi
	User User `gorm:"foreignKey:UserID"`
}

// HashPassword hashes the password
func (ua *UserAuth) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ua.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password matches the hashed password
func (ua *UserAuth) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ua.Password), []byte(password))
	return err == nil
}
