package models

import "time"

// Notification model - Pemberitahuan seperti e-Nofa (notifikasi pajak) atau reminder jatuh tempo.
type Notification struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	UserID       uint      `gorm:"not null;index"`
	Judul        string    `gorm:"type:varchar(255);not null"`
	Isi          string    `gorm:"type:text;not null"`
	Tipe         string    `gorm:"type:varchar(20);not null"` // email/push
	StatusBaca   bool      `gorm:"default:false"`
	TanggalKirim time.Time `gorm:"not null"`
	TaxReturnID  *uint     `gorm:"index"` // optional
	PaymentID    *uint     `gorm:"index"` // optional
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relasi
	User      User       `gorm:"foreignKey:UserID"`
	TaxReturn *TaxReturn `gorm:"foreignKey:TaxReturnID"`
	Payment   *Payment   `gorm:"foreignKey:PaymentID"`
}
