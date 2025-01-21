package models

type Staff struct {
	User           User    `gorm:"foreignKey:UserID" json:"user"`
	JobDescription *string `gorm:"type:text;" json:"jobDescription,omitempty"`
	Salary         float64    `gorm:"not null" json:"salary"`
}
