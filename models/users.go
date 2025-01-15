package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	FirstName string `gorm:"size:256;not null" json:"firstName"`
	LastName  string `gorm:"size:256;not null" json:"lastName"`
}

type User struct {
	gorm.Model
	Email    string `gorm:"size:256;unique; not null" json:"email"`
	Password string `gorm:"size:256;not null" json:"-"`
	Role     string `gorm:"size:25;not null" json:"role"`
	Mobile   string `gorm:"size:15;not null" json:"mobile"`
}
