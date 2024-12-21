package models

type Staff struct {
	User           `gorm:"embedded" json:",inline"`
	JobDescription *string `gorm:"type:text;" json:"jobDescription,omitempty"`
	Salary         uint    `gorm:"not null" json:"salary"`
}
