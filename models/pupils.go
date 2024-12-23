package models

import (
	"gorm.io/gorm"
	"time"
)

type Pupil struct {
	gorm.Model
	FirstName   string    `gorm:"size:256;not null" json:"firstname"`
	LastName    string    `gorm:"size:256;not null" json:"lastname"`
	DOB         time.Time `gorm:"not null" json:"dob"`
	Age         uint      `json:"age"`
	ClassID     uint      `gorm:"not null" json:"classId"`
	ParentID    uint
	Allergies   string  `gorm:"size:256" json:"allergies"`
	MedicalInfo *string `gorm:"type:text" json:"medInfo,omitempty"`
}
