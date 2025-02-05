package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
}

type Role struct {
	gorm.Model
	Name        string       `gorm:"not null;unique" json:"name"`
	Permissions []*Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}
