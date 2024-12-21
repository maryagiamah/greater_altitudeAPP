package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    FirstName string     `gorm:"size:256;not null" json:"firstname"`
    LastName  string     `gorm:"size:256;not null" json:"lastname"`
    Email     *string    `gorm:"size:256;unique" json:"email,omitempty"`
    Password  *string    `gorm:"size:256;" json:"password,omitempty"`
    Role      string     `gorm:"not null" json:"role"`
    Mobile    string     `gorm:"not null" json:"mobile"`
}
