package models

type Staff struct {
	BaseModel
	UserID         uint    `json:"userId"`
	User           User    `gorm:"foreignKey:UserID" json:"user"`
	JobDescription *string `gorm:"type:text;" json:"jobDescription,omitempty"`
	Salary         uint    `gorm:"not null" json:"salary"`
}
