package models

import (
    "time"
    "gorm.io/gorm"
)

type Invoice struct {
    gorm.Model
    PupilID     uint     `gorm:"not null" json:"pupilId"`
    Amount      float64  `gorm:"not null" json:"amount"`
    DueDate     time.Time `gorm:"not null" json:"dueDate"`
    Status      string   `gorm:"size:256;not null" json:"status"`
    Description string   `gorm:"type:text" json:"description,omitempty"`
    Payments []Payment `gorm:"foreignKey:InvoiceID;constraint:OnDelete:CASCADE;" json:"payments,omitempty"`
}

type Payment struct {
    gorm.Model
    InvoiceID   uint     `gorm:"not null" json:"invoice"`
    Amount      float64  `gorm:"not null" json:"amount"`
    PaymentDate time.Time `gorm:"not null" json:"paymentDate"`
    Method      string   `gorm:"size:128;not null" json:"method"`
    Status      string   `gorm:"size:256;not null" json:"status"`
    Reference   string   `gorm:"size:256;not null" json:"reference"`
}

