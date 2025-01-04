package models

import (
	"gorm.io/gorm"
)

type Parent struct {
	gorm.Model
	UserID     uint    `json:"userId"`
	User       User    `gorm:"foreignKey:UserID" json:"user"`
	Address    string  `gorm:"type:text;not null" json:"address"`
	Ward       []Pupil `json:"ward" gorm:"foreignKey:ParentID"`
	Occupation *string `gorm:"size:256" json:"occupation,omitempty"`
}
