package hobbies_commands_test

import (
	"errors"
	"testing"

	cmds "github.com/personalhub/personalhub-back/src/commands/hobbies"
	"github.com/personalhub/personalhub-back/src/internal/testdb"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedReimbursement(t *testing.T, conn *gorm.DB) {
	t.Helper()
	r := hobbies_schemas.Reimbursement{OrderId: "o1", StoreName: "Liga", StoreContactInfo: "sac@liga", StoreOrderId: "L-1", RequestedAt: date1, ReimbursementValue: 1500, ReimbursementStatus: "requested"}
	if err := conn.Create(&r).Error; err != nil {
		t.Fatalf("seeding reimbursement: %v", err)
	}
}

func TestCreateReimbursementCommand(t *testing.T) {
	cmd := cmds.CreateReimbursementCommand{OrderId: "o1", StoreName: "Liga", StoreContactInfo: "sac@liga", StoreOrderId: "L-1", RequestedAt: date1, ReimbursementValue: 1500, ReimbursementStatus: "requested"}

	t.Run("success", func(t *testing.T) {
		testdb.New(t)
		r, err := cmd.Execute()
		if err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if r.OrderId != "o1" || r.StoreName != "Liga" || r.StoreContactInfo != "sac@liga" || r.StoreOrderId != "L-1" ||
			!r.RequestedAt.Equal(date1) || r.ReimbursementValue != 1500 || r.ReimbursementStatus != "requested" {
			t.Errorf("unexpected reimbursement: %+v", r)
		}
	})

	t.Run("duplicate order id", func(t *testing.T) {
		seedReimbursement(t, testdb.New(t))
		if _, err := cmd.Execute(); err == nil {
			t.Fatal("expected an error creating a duplicate reimbursement")
		}
	})
}

func TestUpdateReimbursementCommand(t *testing.T) {
	tests := []struct {
		name  string
		cmd   cmds.UpdateReimbursementCommand
		check func(t *testing.T, r *hobbies_schemas.Reimbursement)
	}{
		{
			name: "all fields",
			cmd: cmds.UpdateReimbursementCommand{StoreName: ptr("Mypcards"), StoreContactInfo: ptr("sac@myp"), StoreOrderId: ptr("M-9"),
				RequestedAt: ptr(date2), ReimbursementValue: ptr(uint(2000)), ReimbursementStatus: ptr("paid")},
			check: func(t *testing.T, r *hobbies_schemas.Reimbursement) {
				if r.StoreName != "Mypcards" || r.StoreContactInfo != "sac@myp" || r.StoreOrderId != "M-9" ||
					!r.RequestedAt.Equal(date2) || r.ReimbursementValue != 2000 || r.ReimbursementStatus != "paid" {
					t.Errorf("fields not updated: %+v", r)
				}
			},
		},
		{
			name: "no fields keeps values",
			cmd:  cmds.UpdateReimbursementCommand{},
			check: func(t *testing.T, r *hobbies_schemas.Reimbursement) {
				if r.StoreName != "Liga" || r.StoreContactInfo != "sac@liga" || r.StoreOrderId != "L-1" ||
					r.ReimbursementValue != 1500 || r.ReimbursementStatus != "requested" {
					t.Errorf("fields changed: %+v", r)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := testdb.New(t)
			seedReimbursement(t, conn)
			tt.cmd.OrderId = "o1"

			r, err := tt.cmd.Execute()
			if err != nil {
				t.Fatalf("Execute: %v", err)
			}
			tt.check(t, r)

			var stored hobbies_schemas.Reimbursement
			conn.First(&stored, "order_id = ?", "o1")
			tt.check(t, &stored)
		})
	}

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if _, err := (cmds.UpdateReimbursementCommand{OrderId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})

	t.Run("database error on save", func(t *testing.T) {
		conn := testdb.New(t)
		seedReimbursement(t, conn)
		testdb.FailWrites(t, conn)
		if _, err := (cmds.UpdateReimbursementCommand{OrderId: "o1", StoreName: ptr("x")}).Execute(); !errors.Is(err, testdb.ErrInjected) {
			t.Fatalf("err = %v, want ErrInjected", err)
		}
	})
}

func TestDeleteReimbursementCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conn := testdb.New(t)
		seedReimbursement(t, conn)
		if err := (cmds.DeleteReimbursementCommand{OrderId: "o1"}).Execute(); err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if err := conn.First(&hobbies_schemas.Reimbursement{}, "order_id = ?", "o1").Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("reimbursement still exists: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if err := (cmds.DeleteReimbursementCommand{OrderId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})
}

func TestGetAndListReimbursementCommands(t *testing.T) {
	seedReimbursement(t, testdb.New(t))

	r, err := (cmds.GetReimbursementCommand{OrderId: "o1"}).Execute()
	if err != nil || r.StoreName != "Liga" {
		t.Fatalf("GetReimbursementCommand = %+v, %v", r, err)
	}

	all, err := (cmds.ListReimbursementsCommand{}).Execute()
	if err != nil || len(all) != 1 {
		t.Fatalf("ListReimbursementsCommand = %d reimbursements, %v", len(all), err)
	}
}
