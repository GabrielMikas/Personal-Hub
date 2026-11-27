package schemas

type Investment struct {
	ID           uint   `gorm:"primarykey"`
	InvestmentId string `gorm:"not null"`
	Type         string `gorm:"not null"`
	Amount       string `gorm:"not null"`
}

func (Investment) TableName() string {
	return "finance.investments"
}
