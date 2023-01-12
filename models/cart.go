package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	KodeProduk int `json:"kode_produk" gorm:"primary_key"`
	NamaProduk string `json:"namaProduk"`
	Kuantitas string `json:"kuatitas"`
}
