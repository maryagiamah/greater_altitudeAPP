package models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	User           User    `gorm:"foreignKey:UserID" json:"user"`
	UserID         uint    `json:"userId"`
	JobDescription *string `gorm:"type:text;" json:"jobDescription,omitempty"`
	Salary         float64 `gorm:"not null" json:"salary"`
}
