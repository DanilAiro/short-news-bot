package models

import (
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	ID         uint  `gorm:"not null;unique" json:"id"`
	Name       string  `gorm:"not null" json:"name"`
	Cost       float64 `gorm:"not null" json:"cost"`
	LastUpdate string  `gorm:"not null" json:"last_update"`
}
