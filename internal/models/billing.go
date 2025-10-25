package models

import "time"

type Billing struct {
	ID                 uint      `gorm:"primaryKey"`
	UserProfileID      uint      `gorm:"not null;index"`
	KodeBilling        string    `gorm:"type:varchar(30);uniqueIndex;not null"`
	KodeAkunPajak      string    `gorm:"type:varchar(10);not null"`
	KodeJenisSetoran   string    `gorm:"type:varchar(10);not null"`
	MasaPajak          int       `gorm:"type:smallint"`
	TahunPajak         int       `gorm:"type:smallint"`
	JumlahSetor        int64     `gorm:"not null"`
	StatusPembayaran   string    `gorm:"type:varchar(50);default:'Belum Dibayar'"`
	TanggalDibuat      time.Time `gorm:"autoCreateTime"`
	TanggalKadaluwarsa time.Time
	NTPN               *string `gorm:"type:varchar(100);uniqueIndex"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
