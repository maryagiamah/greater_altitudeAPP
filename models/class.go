package models

import (
    "gorm.io/gorm"
)

type Class struct {
    gorm.Model
    Name      string `gorm:"size:256;not null" json:"name"`
    AgeGroup  string `gorm:"not null" json:"ageGroup"`
    Pupils    []Pupil `json:"pupils" gorm:"foreignKey:ClassID"`
    ProgramID uint `json:"programId"`
}
