package hobbies_repositories_test

import (
	"errors"
	"testing"
	"time"

	"github.com/personalhub/personalhub-back/src/internal/testdb"
	repo "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

var boughtAt = time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)

func TestCardRepository(t *testing.T) {
	testdb.New(t)

	card := &hobbies_schemas.Card{CardId: "c1", Name: "Pikachu", Code: "SV-025", BoughtAtValue: 1500, BoughtAtDate: boughtAt}
	if err := repo.CreateCard(card); err != nil {
		t.Fatalf("CreateCard: %v", err)
	}
	if err := repo.CreateCard(&hobbies_schemas.Card{CardId: "c2", Name: "Eevee", Code: "SV-133", BoughtAtValue: 900, BoughtAtDate: boughtAt}); err != nil {
		t.Fatalf("CreateCard: %v", err)
	}

	all, err := repo.GetAllCards()
	if err != nil || len(all) != 2 {
		t.Fatalf("GetAllCards = %d cards, %v; want 2, nil", len(all), err)
	}

	got, err := repo.GetCardById("c1")
	if err != nil || got.Name != "Pikachu" || !got.BoughtAtDate.Equal(boughtAt) {
		t.Fatalf("GetCardById = %+v, %v", got, err)
	}

	got.Name = "Raichu"
	if err := repo.UpdateCard(got); err != nil {
		t.Fatalf("UpdateCard: %v", err)
	}
	if got, _ := repo.GetCardById("c1"); got.Name != "Raichu" {
		t.Fatalf("UpdateCard did not persist, name = %q", got.Name)
	}

	if err := repo.DeleteCard("c1"); err != nil {
		t.Fatalf("DeleteCard: %v", err)
	}
	if _, err := repo.GetCardById("c1"); !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("GetCardById after delete: err = %v, want ErrRecordNotFound", err)
	}
}
