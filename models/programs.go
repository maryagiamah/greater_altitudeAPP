package models

import (
	"gorm.io/gorm"
	"time"
)

type Program struct {
	gorm.Model
	Name        string      `gorm:"size:256;not null" json:"name"`
	Description *string     `gorm:"size:256" json:"description,omitempty"`
	StartDate   time.Time   `gorm:"not null" json:"start_date"`
	EndDate     time.Time   `gorm:"not null" json:"end_date"`
	Classes     []*Class    `json:"-"`
	Activities  []*Activity `json:"-"`
}

type Activity struct {
	gorm.Model
	ProgramID   uint    `json:"programId" gorm:"not null"`
	Name        string  `gorm:"not null" json:"name"`
	AgeGroup    string  `gorm:"not null" json:"ageGroup"`
	Description *string `gorm:"not null" json:"description,omitempty"`
}
