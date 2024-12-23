package models

type Staff struct {
	User           `gorm:"embedded"`
	JobDescription *string `gorm:"type:text;" json:"jobDescription,omitempty"`
	Salary         uint    `gorm:"not null" json:"salary"`
}
