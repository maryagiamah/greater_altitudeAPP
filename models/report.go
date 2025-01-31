package models

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	gorm.Model
	PupilID    uint      `gorm:"not null" json:"pupilId"`
	TeacherID  uint      `gorm:"not null" json:"teacherId"`
	Type       string    `gorm:"not null" json:"type"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	Date       time.Time `gorm:"not null" json:"date"`
	Attachment string    `gorm:"type:text" json:"attachment"`

	Pupil  Pupil  `gorm:"foreignKey:PupilID;constraint:OnDelete:CASCADE;" json:"pupil"`
	Teacher Staff `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE;" json:"teacher"`
}
