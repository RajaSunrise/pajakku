package models

import (
	"time"

	"github.com/google/uuid"
)

// User model - Model inti untuk menyimpan data pengguna, termasuk wajib pajak individu atau badan.
type User struct {
	ID                uint      `gorm:"primaryKey;autoIncrement"`
	NIK               uint      `gorm:"type:int;uniqueIndex;not null"`
	NPWP              string    `gorm:"type:varchar(25);uniqueIndex;not null"`
	Nama              string    `gorm:"type:varchar(255);not null"`
	Email             string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Alamat            string    `gorm:"type:text"`
	JenisWajibPajak   string    `gorm:"type:varchar(50);not null"` // individu/badan
	TanggalRegistrasi time.Time `gorm:"autoCreateTime"`
	StatusAktif       bool      `gorm:"default:true"`
	RoleID            uint      `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	// Relasi
	Role          Role           `gorm:"foreignKey:RoleID"`
	TaxReturns    []TaxReturn    `gorm:"foreignKey:UserID"`
	Invoices      []Invoice      `gorm:"foreignKey:UserID"`
	Payments      []Payment      `gorm:"foreignKey:UserID"`
	AuditLogs     []AuditLog     `gorm:"foreignKey:UserID"`
	Notifications []Notification `gorm:"foreignKey:UserID"`
	Attachments   []Attachment   `gorm:"foreignKey:UserID"`
}

// Role model - Mengelola peran akses, seperti admin DJP, wajib pajak, atau auditor.
type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	NamaRole    string `gorm:"type:varchar(50);uniqueIndex;not null"` // e.g., 'wp', 'admin', 'auditor'
	Permissions string `gorm:"type:json"`                             // JSON untuk fitur akses
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relasi
	Users []User `gorm:"foreignKey:RoleID"`
}

// for reset password
type PasswordResetToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `gorm:"index;not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
