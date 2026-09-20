package hobbies_commands

import (
	"time"

	"github.com/google/uuid"
	hobbies_repositories "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

// CreateCardCommand represents the intent to register a new Card.
// CardId is not accepted from the client — it is generated when the
// command executes, since a create request is for a brand new resource.
type CreateCardCommand struct {
	Name          string    `json:"name" binding:"required"`
	Code          string    `json:"code" binding:"required"`
	BoughtAtValue uint      `json:"boughtAtValue" binding:"required"`
	BoughtAtDate  time.Time `json:"boughtAtDate" binding:"required"`
	BinderId      *string   `json:"binderId"`
}

func (cmd CreateCardCommand) Execute() (*hobbies_schemas.Card, error) {
	card := hobbies_schemas.Card{
		CardId:        uuid.New().String(),
		Name:          cmd.Name,
		Code:          cmd.Code,
		BoughtAtValue: cmd.BoughtAtValue,
		BoughtAtDate:  cmd.BoughtAtDate,
		BinderId:      cmd.BinderId,
	}

	if err := hobbies_repositories.CreateCard(&card); err != nil {
		return nil, err
	}
	return &card, nil
}

// UpdateCardCommand represents the intent to update an existing Card.
// CardId is sent through the request (the URL path) to identify which card
// to update; only non-nil fields are applied.
type UpdateCardCommand struct {
	CardId        string     `json:"-"`
	Name          *string    `json:"name"`
	Code          *string    `json:"code"`
	BoughtAtValue *uint      `json:"boughtAtValue"`
	BoughtAtDate  *time.Time `json:"boughtAtDate"`
	BinderId      *string    `json:"binderId"`
}

func (cmd UpdateCardCommand) Execute() (*hobbies_schemas.Card, error) {
	card, err := hobbies_repositories.GetCardById(cmd.CardId)
	if err != nil {
		return nil, err
	}

	if cmd.Name != nil {
		card.Name = *cmd.Name
	}
	if cmd.Code != nil {
		card.Code = *cmd.Code
	}
	if cmd.BoughtAtValue != nil {
		card.BoughtAtValue = *cmd.BoughtAtValue
	}
	if cmd.BoughtAtDate != nil {
		card.BoughtAtDate = *cmd.BoughtAtDate
	}
	if cmd.BinderId != nil {
		card.BinderId = cmd.BinderId
	}

	if err := hobbies_repositories.UpdateCard(card); err != nil {
		return nil, err
	}
	return card, nil
}

// DeleteCardCommand represents the intent to delete a Card. CardId is sent
// through the request (the URL path).
type DeleteCardCommand struct {
	CardId string
}

func (cmd DeleteCardCommand) Execute() error {
	if _, err := hobbies_repositories.GetCardById(cmd.CardId); err != nil {
		return err
	}
	return hobbies_repositories.DeleteCard(cmd.CardId)
}

// GetCardCommand represents the intent to fetch a single Card. CardId is
// sent through the request (the URL path).
type GetCardCommand struct {
	CardId string
}

func (cmd GetCardCommand) Execute() (*hobbies_schemas.Card, error) {
	return hobbies_repositories.GetCardById(cmd.CardId)
}

// ListCardsCommand represents the intent to fetch every Card.
type ListCardsCommand struct{}

func (cmd ListCardsCommand) Execute() ([]hobbies_schemas.Card, error) {
	return hobbies_repositories.GetAllCards()
}
