package router

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewRegistersRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	registered := map[string]bool{}
	for _, r := range New().Routes() {
		registered[r.Method+" "+r.Path] = true
	}

	want := []string{
		http.MethodGet + " /hobbies/cards/",
		http.MethodGet + " /hobbies/cards/:id",
		http.MethodPost + " /hobbies/cards/",
		http.MethodPut + " /hobbies/cards/:id",
		http.MethodDelete + " /hobbies/cards/:id",
		http.MethodGet + " /hobbies/orders/",
		http.MethodGet + " /hobbies/orders/:id",
		http.MethodPost + " /hobbies/orders/",
		http.MethodPut + " /hobbies/orders/:id",
		http.MethodDelete + " /hobbies/orders/:id",
		http.MethodGet + " /hobbies/reimbursements/",
		http.MethodGet + " /hobbies/reimbursements/:orderId",
		http.MethodPost + " /hobbies/reimbursements/",
		http.MethodPut + " /hobbies/reimbursements/:orderId",
		http.MethodDelete + " /hobbies/reimbursements/:orderId",
	}
	for _, route := range want {
		if !registered[route] {
			t.Errorf("route %q is not registered", route)
		}
	}
}
