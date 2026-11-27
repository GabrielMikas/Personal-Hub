package hobbies_controllers_test

import (
	"net/http"
	"testing"

	ctrl "github.com/personalhub/personalhub-back/src/controllers/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedOrder(t *testing.T, conn *gorm.DB) {
	t.Helper()
	if err := conn.Create(&hobbies_schemas.Order{OrderId: "o1", Store: "Liga", StoreOrderId: "L-1", OrderValue: 1500, DeliveryStatus: "pending"}).Error; err != nil {
		t.Fatalf("seeding order: %v", err)
	}
}

func TestOrderController(t *testing.T) {
	h := handlers{ctrl.GetOrders, ctrl.GetOrder, ctrl.CreateOrder, ctrl.UpdateOrder, ctrl.DeleteOrder, "id"}
	valid := `{"store":"Liga","storeOrderId":"L-1","orderedAt":"2026-01-15T00:00:00Z","orderValue":1500,"deliveryStatus":"pending"}`

	runControllerCases(t, "/orders", h, []controllerCase{
		{name: "list", method: http.MethodGet, path: "/orders", setup: []setupFunc{seedOrder}, want: http.StatusOK, check: count(1)},
		{name: "list 500", method: http.MethodGet, path: "/orders", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "get", method: http.MethodGet, path: "/orders/o1", setup: []setupFunc{seedOrder}, want: http.StatusOK, check: field("Store", "Liga")},
		{name: "get 404", method: http.MethodGet, path: "/orders/missing", want: http.StatusNotFound},
		{name: "get 500", method: http.MethodGet, path: "/orders/o1", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "create", method: http.MethodPost, path: "/orders", body: valid, want: http.StatusCreated, check: field("StoreOrderId", "L-1")},
		{name: "create 400 invalid json", method: http.MethodPost, path: "/orders", body: `{`, want: http.StatusBadRequest},
		{name: "create 400 missing field", method: http.MethodPost, path: "/orders", body: `{"store":"Liga"}`, want: http.StatusBadRequest},
		{name: "create 500", method: http.MethodPost, path: "/orders", body: valid, setup: []setupFunc{failWrites}, want: http.StatusInternalServerError},

		{name: "update", method: http.MethodPut, path: "/orders/o1", body: `{"deliveryStatus":"delivered"}`, setup: []setupFunc{seedOrder}, want: http.StatusOK, check: field("DeliveryStatus", "delivered")},
		{name: "update 400", method: http.MethodPut, path: "/orders/o1", body: `not json`, setup: []setupFunc{seedOrder}, want: http.StatusBadRequest},
		{name: "update 404", method: http.MethodPut, path: "/orders/missing", body: `{}`, want: http.StatusNotFound},
		{name: "update 500", method: http.MethodPut, path: "/orders/o1", body: `{"store":"x"}`, setup: []setupFunc{seedOrder, failWrites}, want: http.StatusInternalServerError},

		{name: "delete", method: http.MethodDelete, path: "/orders/o1", setup: []setupFunc{seedOrder}, want: http.StatusNoContent},
		{name: "delete 404", method: http.MethodDelete, path: "/orders/missing", want: http.StatusNotFound},
		{name: "delete 500", method: http.MethodDelete, path: "/orders/o1", setup: []setupFunc{seedOrder, failWrites}, want: http.StatusInternalServerError},
	})
}
