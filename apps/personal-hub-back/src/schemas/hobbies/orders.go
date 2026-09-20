package hobbies_schemas

import (
	"time"
)

type Order struct {
	OrderId        string    `gorm:"primarykey"`
	Store          string    `gorm:"not null"`
	StoreOrderId   string    `gorm:"not null"`
	OrderedAt      time.Time `gorm:"not null"`
	OrderValue     uint      `gorm:"not null"`
	DeliveryStatus string    `gorm:"not null"`
}

func (Order) TableName() string {
	return "hobbies.orders"
}
