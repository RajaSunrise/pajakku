package models

import "time"

// Attachment model - File pendukung seperti bukti pembayaran atau dokumen registrasi.
type Attachment struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	NamaFile     string `gorm:"type:varchar(255);not null"`
	PathURL      string `gorm:"type:varchar(500);not null"`
	TipeMime     string `gorm:"type:varchar(100);not null"`
	Ukuran       int64  `gorm:"not null"`                  // in bytes
	RelatedModel string `gorm:"type:varchar(50);not null"` // e.g., 'TaxReturn'
	UserID       *uint  `gorm:"index"`                     // optional
	TaxReturnID  *uint  `gorm:"index"`                     // optional
	InvoiceID    *uint  `gorm:"index"`                     // optional
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Relasi
	User      *User      `gorm:"foreignKey:UserID"`
	TaxReturn *TaxReturn `gorm:"foreignKey:TaxReturnID"`
	Invoice   *Invoice   `gorm:"foreignKey:InvoiceID"`
}
