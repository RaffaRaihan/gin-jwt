package db

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Nama_Product  	string	`json:"nama_product"`
	Harga			int64	`json:"harga"`
	Stok			int64	`json:"stok"`
}