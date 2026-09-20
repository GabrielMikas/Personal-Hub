package hobbies_repositories

import (
	"github.com/personalhub/personalhub-back/src/config"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

func CreateCard(card *hobbies_schemas.Card) error {
	return config.GetDB().Create(card).Error
}

func GetAllCards() ([]hobbies_schemas.Card, error) {
	var cards []hobbies_schemas.Card
	err := config.GetDB().Find(&cards).Error
	return cards, err
}

func GetCardById(cardId string) (*hobbies_schemas.Card, error) {
	var card hobbies_schemas.Card
	err := config.GetDB().First(&card, "card_id = ?", cardId).Error
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func UpdateCard(card *hobbies_schemas.Card) error {
	return config.GetDB().Save(card).Error
}

func DeleteCard(cardId string) error {
	return config.GetDB().Delete(&hobbies_schemas.Card{}, "card_id = ?", cardId).Error
}
