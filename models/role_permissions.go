package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Action string `gorm:"not null"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"not null;unique"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
