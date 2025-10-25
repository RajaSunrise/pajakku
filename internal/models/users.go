package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Users model
type UserAuth struct {
	ID            uint   `gorm:"primaryKey"`
	UserProfileID *uint  `gorm:"index"`
	FotoProfil    string `gorm:"type:varchar(100)"`
	Email         string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password      string `gorm:"type:varchar(255);not null"`
	StatusAkun    string `gorm:"type:varchar(50);default:'Aktif'"`
	Role          string `gorm:"type:varchar(15); default:'users'"`
	TerakhirLogin *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserProfile struct {
	ID             uint   `gorm:"primaryKey"`
	NPWP           string `gorm:"type:varchar(25);uniqueIndex;not null"`
	NamaWajibPajak string `gorm:"type:varchar(255);not null"`
	TipeWajibPajak string `gorm:"type:varchar(50)"`
	NomorTelepon   string `gorm:"type:varchar(20)"`
	AlamatLengkap  string `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	UserAuth UserAuth `gorm:"foreignKey:UserProfileID;references:ID"`

	Billings   []Billing   `gorm:"foreignKey:UserProfileID"`
	ReportSPTs []ReportSPT `gorm:"foreignKey:UserProfileID"`
}

// for reset password
type PasswordResetToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"index;not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
