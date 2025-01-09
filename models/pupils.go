package models

import (
	"gorm.io/gorm"
	"time"
)

type Pupil struct {
	BaseModel
	DOB         time.Time `gorm:"not null" json:"dob"`
	Age         uint      `json:"age"`
	ClassID     uint      `gorm:"not null" json:"classId"`
	ParentID    uint      `gorm:"not null" json:"parentId"`
	Allergies   string    `gorm:"size:256" json:"allergies"`
	MedicalInfo *string   `gorm:"type:text" json:"medInfo,omitempty"`
}
