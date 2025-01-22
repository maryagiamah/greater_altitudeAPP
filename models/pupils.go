package models

import (
	"time"
)

type Pupil struct {
	BaseModel
	DOB         time.Time `gorm:"not null" json:"dob"`
	ClassID     uint      `gorm:"not null" json:"classId"`
	ParentID    uint      `gorm:"not null" json:"parentId"`
	IsActive    bool      `gorm:"default:true" json:"isActive"`
	Allergies   string    `gorm:"size:256;default:'nil'" json:"allergies"`
	MedicalInfo *string   `gorm:"type:text" json:"medInfo,omitempty"`
}
