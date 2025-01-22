package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string   `gorm:"size:256;not null" json:"name"`
	Pupils    []*Pupil `json:"pupils" gorm:"foreignKey:ClassID"`
	Teachers  []*Staff `gorm:"many2many:class_teachers;" json:"teachers"`
	ProgramID uint     `json:"programId"`
}
