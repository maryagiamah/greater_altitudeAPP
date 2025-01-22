package models

import (
	"gorm.io/gorm"
	"time"
)

type Program struct {
	gorm.Model
	Name        string      `gorm:"size:256;not null" json:"name"`
	Description *string     `gorm:"size:256" json:"description,omitempty"`
	AgeGroup    string      `gorm:"not null" json:"age_group"`
	StartDate   time.Time   `gorm:"not null" json:"start_date"`
	EndDate     time.Time   `gorm:"not null" json:"end_date"`
	Classes     []*Class    `json:"classes"`
	Activities  []*Activity `json:"activities"`
}

type Activity struct {
	gorm.Model
	ProgramID   uint    `json:"curriculumId" gorm:"not null"`
	Name        string  `gorm:"not null" json:"name"`
	AgeGroup    string  `gorm:"not null" json:"age_group"`
	Description *string `gorm:"not null" json:"description,omitempty"`
}
