package hobbies_controllers_test

import (
	"net/http"
	"testing"

	ctrl "github.com/personalhub/personalhub-back/src/controllers/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedCard(t *testing.T, conn *gorm.DB) {
	t.Helper()
	if err := conn.Create(&hobbies_schemas.Card{CardId: "c1", Name: "Pikachu", Code: "SV-025", BoughtAtValue: 1500}).Error; err != nil {
		t.Fatalf("seeding card: %v", err)
	}
}

func TestCardController(t *testing.T) {
	h := handlers{ctrl.GetCards, ctrl.GetCard, ctrl.CreateCard, ctrl.UpdateCard, ctrl.DeleteCard, "id"}
	valid := `{"name":"Pikachu","code":"SV-025","boughtAtValue":1500,"boughtAtDate":"2026-01-15T00:00:00Z"}`

	runControllerCases(t, "/cards", h, []controllerCase{
		{name: "list", method: http.MethodGet, path: "/cards", setup: []setupFunc{seedCard}, want: http.StatusOK, check: count(1)},
		{name: "list 500", method: http.MethodGet, path: "/cards", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "get", method: http.MethodGet, path: "/cards/c1", setup: []setupFunc{seedCard}, want: http.StatusOK, check: field("Name", "Pikachu")},
		{name: "get 404", method: http.MethodGet, path: "/cards/missing", want: http.StatusNotFound},
		{name: "get 500", method: http.MethodGet, path: "/cards/c1", setup: []setupFunc{broken}, want: http.StatusInternalServerError},

		{name: "create", method: http.MethodPost, path: "/cards", body: valid, want: http.StatusCreated, check: field("Code", "SV-025")},
		{name: "create 400 invalid json", method: http.MethodPost, path: "/cards", body: `{`, want: http.StatusBadRequest},
		{name: "create 400 missing field", method: http.MethodPost, path: "/cards", body: `{"name":"Pikachu"}`, want: http.StatusBadRequest},
		{name: "create 500", method: http.MethodPost, path: "/cards", body: valid, setup: []setupFunc{failWrites}, want: http.StatusInternalServerError},

		{name: "update", method: http.MethodPut, path: "/cards/c1", body: `{"name":"Raichu"}`, setup: []setupFunc{seedCard}, want: http.StatusOK, check: field("Name", "Raichu")},
		{name: "update 400", method: http.MethodPut, path: "/cards/c1", body: `not json`, setup: []setupFunc{seedCard}, want: http.StatusBadRequest},
		{name: "update 404", method: http.MethodPut, path: "/cards/missing", body: `{}`, want: http.StatusNotFound},
		{name: "update 500", method: http.MethodPut, path: "/cards/c1", body: `{"name":"Raichu"}`, setup: []setupFunc{seedCard, failWrites}, want: http.StatusInternalServerError},

		{name: "delete", method: http.MethodDelete, path: "/cards/c1", setup: []setupFunc{seedCard}, want: http.StatusNoContent},
		{name: "delete 404", method: http.MethodDelete, path: "/cards/missing", want: http.StatusNotFound},
		{name: "delete 500", method: http.MethodDelete, path: "/cards/c1", setup: []setupFunc{seedCard, failWrites}, want: http.StatusInternalServerError},
	})
}
