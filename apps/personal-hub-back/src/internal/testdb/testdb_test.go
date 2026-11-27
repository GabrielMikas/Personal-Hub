package testdb

import (
	"errors"
	"testing"

	"github.com/personalhub/personalhub-back/src/config"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

func TestNewInstallsMigratedDatabase(t *testing.T) {
	conn := New(t)

	if config.GetDB() != conn {
		t.Fatal("config.GetDB() should return the test database")
	}
	for _, model := range []any{&hobbies_schemas.Card{}, &hobbies_schemas.Order{}, &hobbies_schemas.Reimbursement{}} {
		var count int64
		if err := conn.Model(model).Count(&count).Error; err != nil {
			t.Errorf("table for %T was not migrated: %v", model, err)
		}
	}
}

func TestNewRestoresPreviousDatabase(t *testing.T) {
	outer := New(t)
	t.Run("inner", func(t *testing.T) {
		New(t)
	})
	if config.GetDB() != outer {
		t.Fatal("previous database was not restored after cleanup")
	}
}

func TestFailWrites(t *testing.T) {
	conn := New(t)
	card := hobbies_schemas.Card{CardId: "c1", Name: "n", Code: "c", BoughtAtValue: 1}
	if err := conn.Create(&card).Error; err != nil {
		t.Fatalf("Create before FailWrites: %v", err)
	}

	FailWrites(t, conn)

	if err := conn.Create(&hobbies_schemas.Card{CardId: "c2"}).Error; !errors.Is(err, ErrInjected) {
		t.Errorf("Create error = %v, want ErrInjected", err)
	}
	if err := conn.Save(&card).Error; !errors.Is(err, ErrInjected) {
		t.Errorf("Save error = %v, want ErrInjected", err)
	}
	if err := conn.Delete(&card).Error; !errors.Is(err, ErrInjected) {
		t.Errorf("Delete error = %v, want ErrInjected", err)
	}
	if err := conn.First(&hobbies_schemas.Card{}, "card_id = ?", "c1").Error; err != nil {
		t.Errorf("reads should keep working, got %v", err)
	}
}
