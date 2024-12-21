package models

import (
	"gorm.io/gorm"
	"time"
)

type Enrollment struct {
	gorm.Model
	StudentID uint      `gorm:"default:0" json:"studentId"`
	Status    string    `gorm:"not null" json:"status"`
	ApplyDate time.Time `gorm:"not null" json:"applyDate"`
}
