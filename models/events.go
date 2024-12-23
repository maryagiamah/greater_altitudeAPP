package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null;type:text;" json:"description"`
	Date        time.Time `gorm:"not null" json:"date"`
	Location    string    `gorm:"not null;type:text;" json:"location"`
}
