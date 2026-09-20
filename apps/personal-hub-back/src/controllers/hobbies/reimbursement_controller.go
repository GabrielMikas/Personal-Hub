package hobbies_controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	hobbies_commands "github.com/personalhub/personalhub-back/src/commands/hobbies"
	"gorm.io/gorm"
)

func GetReimbursements(c *gin.Context) {
	cmd := hobbies_commands.ListReimbursementsCommand{}

	reimbursements, err := cmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch reimbursements"})
		return
	}
	c.JSON(http.StatusOK, reimbursements)
}

func GetReimbursement(c *gin.Context) {
	cmd := hobbies_commands.GetReimbursementCommand{OrderId: c.Param("orderId")}

	reimbursement, err := cmd.Execute()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "reimbursement not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch reimbursement"})
		return
	}
	c.JSON(http.StatusOK, reimbursement)
}

func CreateReimbursement(c *gin.Context) {
	var cmd hobbies_commands.CreateReimbursementCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reimbursement, err := cmd.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create reimbursement"})
		return
	}
	c.JSON(http.StatusCreated, reimbursement)
}

func UpdateReimbursement(c *gin.Context) {
	var cmd hobbies_commands.UpdateReimbursementCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.OrderId = c.Param("orderId")

	reimbursement, err := cmd.Execute()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "reimbursement not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update reimbursement"})
		return
	}
	c.JSON(http.StatusOK, reimbursement)
}

func DeleteReimbursement(c *gin.Context) {
	cmd := hobbies_commands.DeleteReimbursementCommand{OrderId: c.Param("orderId")}

	if err := cmd.Execute(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "reimbursement not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete reimbursement"})
		return
	}
	c.Status(http.StatusNoContent)
}
