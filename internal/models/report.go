package models

import "time"

// TaxReturn model - Model untuk pelaporan SPT (Tahunan atau Masa seperti PPh Unifikasi, PPN).
type TaxReturn struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	UserID       uint    `gorm:"not null;index"`
	JenisSPT     string  `gorm:"type:varchar(50);not null"` // tahunan/masa
	PeriodePajak string  `gorm:"type:varchar(50);not null"`
	JumlahPajak  float64 `gorm:"type:decimal(15,2);default:0"`
	Status       string  `gorm:"type:varchar(50);default:'draft'"` // draft/submit/disetujui
	FileSPT      string  `gorm:"type:varchar(255)"`                // path ke PDF/XML
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relasi
	User        User         `gorm:"foreignKey:UserID"`
	Attachments []Attachment `gorm:"foreignKey:TaxReturnID"`
	Invoices    []Invoice    `gorm:"foreignKey:TaxReturnID"`
	Payments    []Payment    `gorm:"foreignKey:ReferensiSPTID"`
}
