package db

import "gorm.io/gorm"

type Link struct {
	*gorm.Model
	Target string
}
