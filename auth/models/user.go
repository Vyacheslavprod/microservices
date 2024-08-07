package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"not null;unique"`
	Hash  string `gorm:"hash" json:"-"`
}