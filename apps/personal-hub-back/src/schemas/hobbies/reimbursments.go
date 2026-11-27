package hobbies_schemas

import (
	"time"
)

type Reimbursement struct {
	OrderId             string    `gorm:"primarykey"`
	StoreName           string    `gorm:"not null"`
	StoreContactInfo    string    `gorm:"not null"`
	StoreOrderId        string    `gorm:"not null"`
	RequestedAt         time.Time `gorm:"not null"`
	ReimbursementValue  uint      `gorm:"not null"`
	ReimbursementStatus string    `gorm:"not null"`
}

func (Reimbursement) TableName() string {
	return "hobbies.reimbursement"
}
