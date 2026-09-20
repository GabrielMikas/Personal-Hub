package hobbies_controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	hobbies_commands "github.com/personalhub/personalhub-back/src/commands/hobbies"
	"gorm.io/gorm"
)

func GetCards(c *gin.Context) {
	cmd := hobbies_commands.ListCardsCommand{}

	cards, err := cmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func GetCard(c *gin.Context) {
	cmd := hobbies_commands.GetCardCommand{CardId: c.Param("id")}

	card, err := cmd.Execute()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch card"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func CreateCard(c *gin.Context) {
	var cmd hobbies_commands.CreateCardCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := cmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create card"})
		return
	}
	c.JSON(http.StatusCreated, card)
}

func UpdateCard(c *gin.Context) {
	var cmd hobbies_commands.UpdateCardCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.CardId = c.Param("id")

	card, err := cmd.Execute()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update card"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func DeleteCard(c *gin.Context) {
	cmd := hobbies_commands.DeleteCardCommand{CardId: c.Param("id")}

	if err := cmd.Execute(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "card not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete card"})
		return
	}
	c.Status(http.StatusNoContent)
}
