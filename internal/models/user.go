package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID string `gorm:"not null;unique" json:"id"`
	User_ID string `gorm:"not null;unique" json:"user_id"`
}