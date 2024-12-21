package models

type Parent struct {
	User       `gorm:"embedded;" json:",inline"`
	Address    string  `gorm:"type:text;not null" json:"address"`
	Ward       []Pupil `json:"ward" gorm:"foreignKey:ParentID"`
	Occupation *string `gorm:"size:256" json:"occupation,omitempty"`
}
