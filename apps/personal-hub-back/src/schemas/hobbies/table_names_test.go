package hobbies_schemas

import "testing"

func TestTableNames(t *testing.T) {
	tests := []struct {
		model interface{ TableName() string }
		want  string
	}{
		{Card{}, "hobbies.cards"},
		{Order{}, "hobbies.orders"},
		{Reimbursement{}, "hobbies.reimbursement"},
	}
	for _, tt := range tests {
		if got := tt.model.TableName(); got != tt.want {
			t.Errorf("%T.TableName() = %q, want %q", tt.model, got, tt.want)
		}
	}
}
