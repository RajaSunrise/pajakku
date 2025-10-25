package models

import "time"

type ReportSPT struct {
	ID            uint      `gorm:"primaryKey"`
	UserProfileID uint      `gorm:"not null;index"` // Foreign Key ke 'user_profiles'
	JenisSPT      string    `gorm:"type:varchar(100);not null"`
	PeriodePajak  string    `gorm:"type:varchar(50);not null"`
	StatusLaporan string    `gorm:"type:varchar(50)"`
	TanggalLapor  time.Time `gorm:"not null"`
	NTTE          string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	FileBPEPath   string    `gorm:"type:varchar(255)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
