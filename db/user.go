package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama	 string	`json:"nama"`
	Email    string `gorm:"unique"`
	Password string
	Telepon  int 	`json:"telepon"`
}