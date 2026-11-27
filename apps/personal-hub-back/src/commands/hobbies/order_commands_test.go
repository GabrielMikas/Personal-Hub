package hobbies_commands_test

import (
	"errors"
	"testing"

	cmds "github.com/personalhub/personalhub-back/src/commands/hobbies"
	"github.com/personalhub/personalhub-back/src/internal/testdb"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedOrder(t *testing.T, conn *gorm.DB) hobbies_schemas.Order {
	t.Helper()
	order := hobbies_schemas.Order{OrderId: "o1", Store: "Liga", StoreOrderId: "L-1", OrderedAt: date1, OrderValue: 1500, DeliveryStatus: "pending"}
	if err := conn.Create(&order).Error; err != nil {
		t.Fatalf("seeding order: %v", err)
	}
	return order
}

func TestCreateOrderCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testdb.New(t)
		cmd := cmds.CreateOrderCommand{Store: "Liga", StoreOrderId: "L-1", OrderedAt: date1, OrderValue: 1500, DeliveryStatus: "pending"}

		order, err := cmd.Execute()
		if err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if order.OrderId == "" {
			t.Error("OrderId should be generated")
		}
		if order.Store != "Liga" || order.StoreOrderId != "L-1" || order.OrderValue != 1500 || !order.OrderedAt.Equal(date1) || order.DeliveryStatus != "pending" {
			t.Errorf("unexpected order: %+v", order)
		}
	})

	t.Run("database error", func(t *testing.T) {
		testdb.FailWrites(t, testdb.New(t))
		if _, err := (cmds.CreateOrderCommand{Store: "s", StoreOrderId: "x", OrderValue: 1, DeliveryStatus: "d"}).Execute(); !errors.Is(err, testdb.ErrInjected) {
			t.Fatalf("err = %v, want ErrInjected", err)
		}
	})
}

func TestUpdateOrderCommand(t *testing.T) {
	tests := []struct {
		name  string
		cmd   cmds.UpdateOrderCommand
		check func(t *testing.T, c *hobbies_schemas.Order)
	}{
		{
			name: "all fields",
			cmd:  cmds.UpdateOrderCommand{Store: ptr("Mypcards"), StoreOrderId: ptr("M-9"), OrderedAt: ptr(date2), OrderValue: ptr(uint(2000)), DeliveryStatus: ptr("delivered")},
			check: func(t *testing.T, c *hobbies_schemas.Order) {
				if c.Store != "Mypcards" || c.StoreOrderId != "M-9" || c.OrderValue != 2000 || !c.OrderedAt.Equal(date2) || c.DeliveryStatus != "delivered" {
					t.Errorf("fields not updated: %+v", c)
				}
			},
		},
		{
			name: "no fields keeps values",
			cmd:  cmds.UpdateOrderCommand{},
			check: func(t *testing.T, c *hobbies_schemas.Order) {
				if c.Store != "Liga" || c.StoreOrderId != "L-1" || c.OrderValue != 1500 || c.DeliveryStatus != "pending" {
					t.Errorf("fields changed: %+v", c)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := testdb.New(t)
			seedOrder(t, conn)
			tt.cmd.OrderId = "o1"

			order, err := tt.cmd.Execute()
			if err != nil {
				t.Fatalf("Execute: %v", err)
			}
			tt.check(t, order)

			var stored hobbies_schemas.Order
			conn.First(&stored, "order_id = ?", "o1")
			tt.check(t, &stored)
		})
	}

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if _, err := (cmds.UpdateOrderCommand{OrderId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})

	t.Run("database error on save", func(t *testing.T) {
		conn := testdb.New(t)
		seedOrder(t, conn)
		testdb.FailWrites(t, conn)
		if _, err := (cmds.UpdateOrderCommand{OrderId: "o1", Store: ptr("x")}).Execute(); !errors.Is(err, testdb.ErrInjected) {
			t.Fatalf("err = %v, want ErrInjected", err)
		}
	})
}

func TestDeleteOrderCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conn := testdb.New(t)
		seedOrder(t, conn)
		if err := (cmds.DeleteOrderCommand{OrderId: "o1"}).Execute(); err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if err := conn.First(&hobbies_schemas.Order{}, "order_id = ?", "o1").Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("order still exists: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if err := (cmds.DeleteOrderCommand{OrderId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})
}

func TestGetAndListOrderCommands(t *testing.T) {
	conn := testdb.New(t)
	seedOrder(t, conn)

	order, err := (cmds.GetOrderCommand{OrderId: "o1"}).Execute()
	if err != nil || order.Store != "Liga" {
		t.Fatalf("GetOrderCommand = %+v, %v", order, err)
	}

	orders, err := (cmds.ListOrdersCommand{}).Execute()
	if err != nil || len(orders) != 1 {
		t.Fatalf("ListOrdersCommand = %d orders, %v", len(orders), err)
	}
}
