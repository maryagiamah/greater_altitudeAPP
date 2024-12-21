package models

import (
    "gorm.io/gorm"
    "time"
)

type Program struct {
    gorm.Model
    Name        string    `gorm:"size:256;not null" json:"name"`
    Description *string   `gorm:"size:256" json:"description,omitempty"`
    AgeGroup    string    `gorm:"not null" json:"age_group"`
    StartDate   time.Time `gorm:"not null" json:"start_date"`
    EndDate     time.Time `gorm:"not null" json:"end_date"`
    Classes     []Class   `json:"classes" gorm:"foreignKey:ProgramID"`
}
