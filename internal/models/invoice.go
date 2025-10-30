package models

import "time"

// Invoice model - Penyimpanan faktur pajak masuk/keluar, termasuk e-Faktur PPN.
type Invoice struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	NomorFaktur      string    `gorm:"type:varchar(50);uniqueIndex;not null"`
	UserID           uint      `gorm:"not null;index"` // penerbit/penerima
	TanggalTransaksi time.Time `gorm:"not null"`
	Jumlah           float64   `gorm:"type:decimal(15,2);not null"`
	Jenis            string    `gorm:"type:varchar(20);not null"` // masuk/keluar
	StatusVerifikasi string    `gorm:"type:varchar(50);default:'pending'"`
	TaxReturnID      *uint     `gorm:"index"` // optional
	CreatedAt        time.Time
	UpdatedAt        time.Time

	// Relasi
	User        User         `gorm:"foreignKey:UserID"`
	TaxReturn   *TaxReturn   `gorm:"foreignKey:TaxReturnID"`
	Attachments []Attachment `gorm:"foreignKey:InvoiceID"`
}
