package models

type Parent struct {
	BaseModel
	UserID     uint    `json:"userId"`
	User       User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	Address    string  `gorm:"type:text;not null" json:"address"`
	Ward       []Pupil `json:"ward" gorm:"foreignKey:ParentID"`
	Occupation *string `gorm:"size:256" json:"occupation,omitempty"`
}
