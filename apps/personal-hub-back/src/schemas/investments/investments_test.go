package schemas

import "testing"

func TestInvestmentTableName(t *testing.T) {
	if got := (Investment{}).TableName(); got != "finance.investments" {
		t.Errorf("TableName() = %q, want %q", got, "finance.investments")
	}
}
