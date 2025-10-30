package models

import "time"

// AuditLog model - Riwayat aktivitas untuk compliance dan troubleshooting.
type AuditLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    uint      `gorm:"not null;index"`
	Aksi      string    `gorm:"type:varchar(100);not null"` // e.g., 'submit_spt', 'update_data'
	Timestamp time.Time `gorm:"autoCreateTime"`
	IPAddress string    `gorm:"type:varchar(45)"` // IPv4/IPv6
	OldValue  string    `gorm:"type:json"`        // JSON
	NewValue  string    `gorm:"type:json"`        // JSON
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relasi
	User User `gorm:"foreignKey:UserID"`
}
