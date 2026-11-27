package hobbies_commands

import (
	"time"

	"github.com/google/uuid"
	hobbies_repositories "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

// CreateOrderCommand represents the intent to register a new Order.
// OrderId is not accepted from the client — it is generated when the
// command executes, since a create request is for a brand new resource.
type CreateOrderCommand struct {
	Store          string    `json:"store" binding:"required"`
	StoreOrderId   string    `json:"storeOrderId" binding:"required"`
	OrderedAt      time.Time `json:"orderedAt" binding:"required"`
	OrderValue     uint      `json:"orderValue" binding:"required"`
	DeliveryStatus string    `json:"deliveryStatus" binding:"required"`
}

func (cmd CreateOrderCommand) Execute() (*hobbies_schemas.Order, error) {
	order := hobbies_schemas.Order{
		OrderId:        uuid.New().String(),
		Store:          cmd.Store,
		StoreOrderId:   cmd.StoreOrderId,
		OrderedAt:      cmd.OrderedAt,
		OrderValue:     cmd.OrderValue,
		DeliveryStatus: cmd.DeliveryStatus,
	}

	if err := hobbies_repositories.CreateOrder(&order); err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrderCommand represents the intent to update an existing Order.
// OrderId is sent through the request (the URL path) to identify which
// order to update; only non-nil fields are applied.
type UpdateOrderCommand struct {
	OrderId        string     `json:"-"`
	Store          *string    `json:"store"`
	StoreOrderId   *string    `json:"storeOrderId"`
	OrderedAt      *time.Time `json:"orderedAt"`
	OrderValue     *uint      `json:"orderValue"`
	DeliveryStatus *string    `json:"deliveryStatus"`
}

func (cmd UpdateOrderCommand) Execute() (*hobbies_schemas.Order, error) {
	order, err := hobbies_repositories.GetOrderById(cmd.OrderId)
	if err != nil {
		return nil, err
	}

	if cmd.Store != nil {
		order.Store = *cmd.Store
	}
	if cmd.StoreOrderId != nil {
		order.StoreOrderId = *cmd.StoreOrderId
	}
	if cmd.OrderedAt != nil {
		order.OrderedAt = *cmd.OrderedAt
	}
	if cmd.OrderValue != nil {
		order.OrderValue = *cmd.OrderValue
	}
	if cmd.DeliveryStatus != nil {
		order.DeliveryStatus = *cmd.DeliveryStatus
	}

	if err := hobbies_repositories.UpdateOrder(order); err != nil {
		return nil, err
	}
	return order, nil
}

// DeleteOrderCommand represents the intent to delete an Order. OrderId is
// sent through the request (the URL path).
type DeleteOrderCommand struct {
	OrderId string
}

func (cmd DeleteOrderCommand) Execute() error {
	if _, err := hobbies_repositories.GetOrderById(cmd.OrderId); err != nil {
		return err
	}
	return hobbies_repositories.DeleteOrder(cmd.OrderId)
}

// GetOrderCommand represents the intent to fetch a single Order. OrderId is
// sent through the request (the URL path).
type GetOrderCommand struct {
	OrderId string
}

func (cmd GetOrderCommand) Execute() (*hobbies_schemas.Order, error) {
	return hobbies_repositories.GetOrderById(cmd.OrderId)
}

// ListOrdersCommand represents the intent to fetch every Order.
type ListOrdersCommand struct{}

func (cmd ListOrdersCommand) Execute() ([]hobbies_schemas.Order, error) {
	return hobbies_repositories.GetAllOrders()
}
