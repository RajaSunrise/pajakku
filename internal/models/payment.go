package models

import "time"

// Payment model - Transaksi pembayaran pajak, termasuk e-Bupot dan bukti potong.
type Payment struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	UserID           uint      `gorm:"not null;index"`
	JumlahBayar      float64   `gorm:"type:decimal(15,2);not null"`
	MetodePembayaran string    `gorm:"type:varchar(50);not null"` // transfer/virtual account
	TanggalBayar     time.Time `gorm:"not null"`
	ReferensiSPTID   *uint     `gorm:"index"`                              // optional
	Status           string    `gorm:"type:varchar(50);default:'pending'"` // pending/sukses
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// Relasi
	User      User       `gorm:"foreignKey:UserID"`
	TaxReturn *TaxReturn `gorm:"foreignKey:ReferensiSPTID"`
}
