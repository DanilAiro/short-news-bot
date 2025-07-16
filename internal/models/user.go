package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"not null;unique" json:"id"`
	User_ID int64 `gorm:"not null;unique" json:"user_id"`
}