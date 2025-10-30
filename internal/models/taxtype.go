package models

import "time"

// TaxType model - Katalog jenis pajak yang didukung (PPh, PPN, PPN Pemungut, dll.).
type TaxType struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	KodePajak    string  `gorm:"type:varchar(20);uniqueIndex;not null"`
	Nama         string  `gorm:"type:varchar(100);not null"`
	TarifDefault float64 `gorm:"type:decimal(5,2);default:0"`
	Deskripsi    string  `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Many-to-Many relations
	TaxReturns []TaxReturn `gorm:"many2many:taxreturn_taxtypes;"`
	Invoices   []Invoice   `gorm:"many2many:invoice_taxtypes;"`
}
