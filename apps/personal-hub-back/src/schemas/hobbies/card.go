package hobbies_schemas

import (
	"time"
)

type Card struct {
	CardId        string    `gorm:"primarykey"`
	Name          string    `gorm:"not null"`
	Code          string    `gorm:"not null"`
	BoughtAtValue uint      `gorm:"not null"`
	BoughtAtDate  time.Time `gorm:"not null"`
	BinderId      *string
}

func (Card) TableName() string {
	return "hobbies.cards"
}
