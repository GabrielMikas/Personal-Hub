package hobbies_commands_test

import (
	"errors"
	"testing"

	cmds "github.com/personalhub/personalhub-back/src/commands/hobbies"
	"github.com/personalhub/personalhub-back/src/internal/testdb"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
	"gorm.io/gorm"
)

func seedCard(t *testing.T, conn *gorm.DB) hobbies_schemas.Card {
	t.Helper()
	card := hobbies_schemas.Card{CardId: "c1", Name: "Pikachu", Code: "SV-025", BoughtAtValue: 1500, BoughtAtDate: date1}
	if err := conn.Create(&card).Error; err != nil {
		t.Fatalf("seeding card: %v", err)
	}
	return card
}

func TestCreateCardCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		testdb.New(t)
		cmd := cmds.CreateCardCommand{Name: "Pikachu", Code: "SV-025", BoughtAtValue: 1500, BoughtAtDate: date1, BinderId: ptr("b1")}

		card, err := cmd.Execute()
		if err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if card.CardId == "" {
			t.Error("CardId should be generated")
		}
		if card.Name != "Pikachu" || card.Code != "SV-025" || card.BoughtAtValue != 1500 || !card.BoughtAtDate.Equal(date1) || *card.BinderId != "b1" {
			t.Errorf("unexpected card: %+v", card)
		}
	})

	t.Run("database error", func(t *testing.T) {
		testdb.FailWrites(t, testdb.New(t))
		if _, err := (cmds.CreateCardCommand{Name: "n", Code: "c", BoughtAtValue: 1}).Execute(); !errors.Is(err, testdb.ErrInjected) {
			t.Fatalf("err = %v, want ErrInjected", err)
		}
	})
}

func TestUpdateCardCommand(t *testing.T) {
	tests := []struct {
		name  string
		cmd   cmds.UpdateCardCommand
		check func(t *testing.T, c *hobbies_schemas.Card)
	}{
		{
			name: "all fields",
			cmd:  cmds.UpdateCardCommand{Name: ptr("Raichu"), Code: ptr("SV-026"), BoughtAtValue: ptr(uint(2000)), BoughtAtDate: ptr(date2), BinderId: ptr("b2")},
			check: func(t *testing.T, c *hobbies_schemas.Card) {
				if c.Name != "Raichu" || c.Code != "SV-026" || c.BoughtAtValue != 2000 || !c.BoughtAtDate.Equal(date2) || *c.BinderId != "b2" {
					t.Errorf("fields not updated: %+v", c)
				}
			},
		},
		{
			name: "no fields keeps values",
			cmd:  cmds.UpdateCardCommand{},
			check: func(t *testing.T, c *hobbies_schemas.Card) {
				if c.Name != "Pikachu" || c.Code != "SV-025" || c.BoughtAtValue != 1500 || c.BinderId != nil {
					t.Errorf("fields changed: %+v", c)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := testdb.New(t)
			seedCard(t, conn)
			tt.cmd.CardId = "c1"

			card, err := tt.cmd.Execute()
			if err != nil {
				t.Fatalf("Execute: %v", err)
			}
			tt.check(t, card)

			var stored hobbies_schemas.Card
			conn.First(&stored, "card_id = ?", "c1")
			tt.check(t, &stored)
		})
	}

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if _, err := (cmds.UpdateCardCommand{CardId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})

	t.Run("database error on save", func(t *testing.T) {
		conn := testdb.New(t)
		seedCard(t, conn)
		testdb.FailWrites(t, conn)
		if _, err := (cmds.UpdateCardCommand{CardId: "c1", Name: ptr("x")}).Execute(); !errors.Is(err, testdb.ErrInjected) {
			t.Fatalf("err = %v, want ErrInjected", err)
		}
	})
}

func TestDeleteCardCommand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		conn := testdb.New(t)
		seedCard(t, conn)
		if err := (cmds.DeleteCardCommand{CardId: "c1"}).Execute(); err != nil {
			t.Fatalf("Execute: %v", err)
		}
		if err := conn.First(&hobbies_schemas.Card{}, "card_id = ?", "c1").Error; !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("card still exists: %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		testdb.New(t)
		if err := (cmds.DeleteCardCommand{CardId: "missing"}).Execute(); !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatalf("err = %v, want ErrRecordNotFound", err)
		}
	})
}

func TestGetAndListCardCommands(t *testing.T) {
	conn := testdb.New(t)
	seedCard(t, conn)

	card, err := (cmds.GetCardCommand{CardId: "c1"}).Execute()
	if err != nil || card.Name != "Pikachu" {
		t.Fatalf("GetCardCommand = %+v, %v", card, err)
	}

	cards, err := (cmds.ListCardsCommand{}).Execute()
	if err != nil || len(cards) != 1 {
		t.Fatalf("ListCardsCommand = %d cards, %v", len(cards), err)
	}
}
