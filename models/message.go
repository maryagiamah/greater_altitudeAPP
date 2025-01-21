type Message struct {
	SenderID     uint `gorm:"not null" json:"senderId"`
	ReceiverID   uint `gorm:"not null" json:"receiverId"`
	Content      string `json:"content" gorm:"not null"`
	IsRead       bool   `json:"isRead" gorm:"not null"`
	Attachment   *string `json:"attachment,omitempty"`
	Priority     string  `gorm:"priority" json:"priority"`

	Sender   User `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE;" json:"-"`
	Receiver User `gorm:"foreignKey:ReceiverID;constraint:OnDelete:CASCADE;" json:"-"`
}
