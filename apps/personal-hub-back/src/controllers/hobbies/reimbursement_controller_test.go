package hobbies_controllers_test

import (
	"net/http"
	"testing"

	ctrl "github.com/personalhub/personalhub-back/src/controllers/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedReimbursement(t *testing.T, conn *gorm.DB) {
	t.Helper()
	r := hobbies_schemas.Reimbursement{OrderId: "o1", StoreName: "Liga", StoreContactInfo: "sac@liga", StoreOrderId: "L-1", ReimbursementValue: 1500, ReimbursementStatus: "requested"}
	if err := conn.Create(&r).Error; err != nil {
		t.Fatalf("seeding reimbursement: %v", err)
	}
}

func TestReimbursementController(t *testing.T) {
	h := handlers{ctrl.GetReimbursements, ctrl.GetReimbursement, ctrl.CreateReimbursement, ctrl.UpdateReimbursement, ctrl.DeleteReimbursement, "orderId"}
	valid := `{"orderId":"o2","storeName":"Liga","storeContactInfo":"sac@liga","storeOrderId":"L-2","requestedAt":"2026-01-15T00:00:00Z","reimbursementValue":1500,"reimbursementStatus":"requested"}`
	duplicate := `{"orderId":"o1","storeName":"Liga","storeContactInfo":"sac@liga","storeOrderId":"L-1","requestedAt":"2026-01-15T00:00:00Z","reimbursementValue":1500,"reimbursementStatus":"requested"}`

	runControllerCases(t, "/reimbursements", h, []controllerCase{
		{name: "list", method: http.MethodGet, path: "/reimbursements", setup: []setupFunc{seedReimbursement}, want: http.StatusOK, check: count(1)},
		{name: "list 500", method: http.MethodGet, path: "/reimbursements", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "get", method: http.MethodGet, path: "/reimbursements/o1", setup: []setupFunc{seedReimbursement}, want: http.StatusOK, check: field("StoreName", "Liga")},
		{name: "get 404", method: http.MethodGet, path: "/reimbursements/missing", want: http.StatusNotFound},
		{name: "get 500", method: http.MethodGet, path: "/reimbursements/o1", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "create", method: http.MethodPost, path: "/reimbursements", body: valid, want: http.StatusCreated, check: field("OrderId", "o2")},
		{name: "create 400 invalid json", method: http.MethodPost, path: "/reimbursements", body: `{`, want: http.StatusBadRequest},
		{name: "create 400 missing field", method: http.MethodPost, path: "/reimbursements", body: `{"orderId":"o2"}`, want: http.StatusBadRequest},
		{name: "create 500 duplicate", method: http.MethodPost, path: "/reimbursements", body: duplicate, setup: []setupFunc{seedReimbursement}, want: http.StatusInternalServerError},

		{name: "update", method: http.MethodPut, path: "/reimbursements/o1", body: `{"reimbursementStatus":"paid"}`, setup: []setupFunc{seedReimbursement}, want: http.StatusOK, check: field("ReimbursementStatus", "paid")},
		{name: "update 400", method: http.MethodPut, path: "/reimbursements/o1", body: `not json`, setup: []setupFunc{seedReimbursement}, want: http.StatusBadRequest},
		{name: "update 404", method: http.MethodPut, path: "/reimbursements/missing", body: `{}`, want: http.StatusNotFound},
		{name: "update 500", method: http.MethodPut, path: "/reimbursements/o1", body: `{"storeName":"x"}`, setup: []setupFunc{seedReimbursement, failWrites}, want: http.StatusInternalServerError},

		{name: "delete", method: http.MethodDelete, path: "/reimbursements/o1", setup: []setupFunc{seedReimbursement}, want: http.StatusNoContent},
		{name: "delete 404", method: http.MethodDelete, path: "/reimbursements/missing", want: http.StatusNotFound},
		{name: "delete 500", method: http.MethodDelete, path: "/reimbursements/o1", setup: []setupFunc{seedReimbursement, failWrites}, want: http.StatusInternalServerError},
	})
}
