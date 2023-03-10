package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null" json:"name"`
	Telephone string `gorm:"type:varchar(11);not null" json:"telephone"`
	Password  string `gorm:"size:255;not null" json:"password"`
}
