package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserEmail string `gorm:"unique"`
	Password  string `gorm:"not null"`
	IsBan     bool   `gorm:"default:false"`
}
