package models

import (
	"gorm.io/gorm"
	"time"
)

type News struct {
	gorm.Model
	Title    string    `json:"title" gorm:"size:256;not null"`
	Content  string    `json:"content" gorm:"type:text;not null"`
	Date     time.Time `json:"date" gorm:"not null"`
	AuthorID uint      `json:"authorId" gorm:"not null"`
}
