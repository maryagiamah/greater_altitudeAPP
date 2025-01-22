package models

import (
	"gorm.io/gorm"
)

type Parent struct {
	gorm.Model
	User       User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	UserID     uint    `json:"userId"`
	Address    string  `gorm:"type:text;not null" json:"address"`
	Ward       []Pupil `json:"ward" gorm:"foreignKey:ParentID"`
	Occupation *string `gorm:"size:256" json:"occupation,omitempty"`
}
