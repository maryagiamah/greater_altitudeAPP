package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string   `gorm:"size:256;not null" json:"name"`
	Pupils    []*Pupil `json:"-" gorm:"foreignKey:ClassID"`
	Teachers  []*Staff `gorm:"many2many:class_teachers;" json:"-"`
	ProgramID uint     `json:"programId"`
}
