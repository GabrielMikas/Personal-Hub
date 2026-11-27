package hobbies_repositories

import (
	"github.com/personalhub/personalhub-back/src/config"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

func CreateOrder(order *hobbies_schemas.Order) error {
	return config.GetDB().Create(order).Error
}

func GetAllOrders() ([]hobbies_schemas.Order, error) {
	var orders []hobbies_schemas.Order
	err := config.GetDB().Find(&orders).Error
	return orders, err
}

func GetOrderById(orderId string) (*hobbies_schemas.Order, error) {
	var order hobbies_schemas.Order
	err := config.GetDB().First(&order, "order_id = ?", orderId).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(order *hobbies_schemas.Order) error {
	return config.GetDB().Save(order).Error
}

func DeleteOrder(orderId string) error {
	return config.GetDB().Delete(&hobbies_schemas.Order{}, "order_id = ?", orderId).Error
}
