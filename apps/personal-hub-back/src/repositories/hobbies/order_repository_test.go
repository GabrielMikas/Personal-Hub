package hobbies_repositories_test

import (
	"errors"
	"testing"

	"github.com/personalhub/personalhub-back/src/internal/testdb"
	repo "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func TestOrderRepository(t *testing.T) {
	testdb.New(t)

	order := &hobbies_schemas.Order{OrderId: "o1", Store: "Liga", StoreOrderId: "L-1", OrderedAt: boughtAt, OrderValue: 5000, DeliveryStatus: "pending"}
	if err := repo.CreateOrder(order); err != nil {
		t.Fatalf("CreateOrder: %v", err)
	}
	if err := repo.CreateOrder(&hobbies_schemas.Order{OrderId: "o2", Store: "Liga", StoreOrderId: "L-2", OrderedAt: boughtAt, OrderValue: 100, DeliveryStatus: "delivered"}); err != nil {
		t.Fatalf("CreateOrder: %v", err)
	}

	all, err := repo.GetAllOrders()
	if err != nil || len(all) != 2 {
		t.Fatalf("GetAllOrders = %d orders, %v; want 2, nil", len(all), err)
	}

	got, err := repo.GetOrderById("o1")
	if err != nil || got.Store != "Liga" || !got.OrderedAt.Equal(boughtAt) {
		t.Fatalf("GetOrderById = %+v, %v", got, err)
	}

	got.DeliveryStatus = "delivered"
	if err := repo.UpdateOrder(got); err != nil {
		t.Fatalf("UpdateOrder: %v", err)
	}
	if got, _ := repo.GetOrderById("o1"); got.DeliveryStatus != "delivered" {
		t.Fatalf("UpdateOrder did not persist, status = %q", got.DeliveryStatus)
	}

	if err := repo.DeleteOrder("o1"); err != nil {
		t.Fatalf("DeleteOrder: %v", err)
	}
	if _, err := repo.GetOrderById("o1"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("GetOrderById after delete: err = %v, want ErrRecordNotFound", err)
	}
}
