package hobbies_controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/personalhub/personalhub-back/src/internal/testdb"
	"gorm.io/gorm"
)

func init() { gin.SetMode(gin.TestMode) }

// handlers holds the five CRUD handlers of a resource.
type handlers struct {
	list, get, create, update, delete gin.HandlerFunc
	param                             string // path parameter name, e.g. "id"
}

func newRouter(base string, h handlers) *gin.Engine {
	r := gin.New()
	r.GET(base, h.list)
	r.GET(base+"/:"+h.param, h.get)
	r.POST(base, h.create)
	r.PUT(base+"/:"+h.param, h.update)
	r.DELETE(base+"/:"+h.param, h.delete)
	return r
}

type setupFunc func(t *testing.T, conn *gorm.DB)

// broken closes the database so every query fails with a non-404 error.
func broken(t *testing.T, conn *gorm.DB) {
	sqlDB, _ := conn.DB()
	sqlDB.Close()
}

type controllerCase struct {
	name   string
	method string
	path   string
	body   string
	setup  []setupFunc
	want   int
	check  func(t *testing.T, rec *httptest.ResponseRecorder) // optional
}

func runControllerCases(t *testing.T, base string, h handlers, cases []controllerCase) {
	t.Helper()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conn := testdb.New(t)
			for _, setup := range tc.setup {
				setup(t, conn)
			}

			req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			newRouter(base, h).ServeHTTP(rec, req)

			if rec.Code != tc.want {
				t.Fatalf("%s %s = %d, want %d; body: %s", tc.method, tc.path, rec.Code, tc.want, rec.Body)
			}
			if tc.want >= http.StatusBadRequest {
				var body map[string]any
				if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil || body["error"] == nil {
					t.Errorf("error response should have an \"error\" field, got %s", rec.Body)
				}
			}
			if tc.check != nil {
				tc.check(t, rec)
			}
		})
	}
}

// field checks a string field of a JSON object response.
func field(key, want string) func(t *testing.T, rec *httptest.ResponseRecorder) {
	return func(t *testing.T, rec *httptest.ResponseRecorder) {
		var body map[string]any
		if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
			t.Fatalf("decoding body %q: %v", rec.Body, err)
		}
		if body[key] != want {
			t.Errorf("%s = %v, want %q", key, body[key], want)
		}
	}
}

// count checks the number of items of a JSON array response.
func count(want int) func(t *testing.T, rec *httptest.ResponseRecorder) {
	return func(t *testing.T, rec *httptest.ResponseRecorder) {
		var items []any
		if err := json.Unmarshal(rec.Body.Bytes(), &items); err != nil {
			t.Fatalf("decoding list %q: %v", rec.Body, err)
		}
		if len(items) != want {
			t.Errorf("got %d items, want %d", len(items), want)
		}
	}
}

// failWrites makes creates, updates and deletes fail while reads keep working.
func failWrites(t *testing.T, conn *gorm.DB) { testdb.FailWrites(t, conn) }
