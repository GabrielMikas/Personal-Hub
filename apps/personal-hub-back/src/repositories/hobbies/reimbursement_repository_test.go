package hobbies_repositories_test

import (
	"errors"
	"testing"

	"github.com/personalhub/personalhub-back/src/internal/testdb"
	repo "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func TestReimbursementRepository(t *testing.T) {
	testdb.New(t)

	r := &hobbies_schemas.Reimbursement{OrderId: "o1", StoreName: "Liga", StoreContactInfo: "sac@liga", StoreOrderId: "L-1", RequestedAt: boughtAt, ReimbursementValue: 5000, ReimbursementStatus: "requested"}
	if err := repo.CreateReimbursement(r); err != nil {
		t.Fatalf("CreateReimbursement: %v", err)
	}
	if err := repo.CreateReimbursement(&hobbies_schemas.Reimbursement{OrderId: "o2", StoreName: "Liga", StoreContactInfo: "sac@liga", StoreOrderId: "L-2", RequestedAt: boughtAt, ReimbursementValue: 100, ReimbursementStatus: "paid"}); err != nil {
		t.Fatalf("CreateReimbursement: %v", err)
	}

	all, err := repo.GetAllReimbursements()
	if err != nil || len(all) != 2 {
		t.Fatalf("GetAllReimbursements = %d reimbursements, %v; want 2, nil", len(all), err)
	}

	got, err := repo.GetReimbursementById("o1")
	if err != nil || got.StoreName != "Liga" || !got.RequestedAt.Equal(boughtAt) {
		t.Fatalf("GetReimbursementById = %+v, %v", got, err)
	}

	got.ReimbursementStatus = "paid"
	if err := repo.UpdateReimbursement(got); err != nil {
		t.Fatalf("UpdateReimbursement: %v", err)
	}
	if got, _ := repo.GetReimbursementById("o1"); got.ReimbursementStatus != "paid" {
		t.Fatalf("UpdateReimbursement did not persist, status = %q", got.ReimbursementStatus)
	}

	if err := repo.DeleteReimbursement("o1"); err != nil {
		t.Fatalf("DeleteReimbursement: %v", err)
	}
	if _, err := repo.GetReimbursementById("o1"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("GetReimbursementById after delete: err = %v, want ErrRecordNotFound", err)
	}
}
