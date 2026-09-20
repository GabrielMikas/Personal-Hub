package hobbies_repositories

import (
	"github.com/personalhub/personalhub-back/src/config"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

func CreateReimbursement(reimbursement *hobbies_schemas.Reimbursement) error {
	return config.GetDB().Create(reimbursement).Error
}

func GetAllReimbursements() ([]hobbies_schemas.Reimbursement, error) {
	var reimbursements []hobbies_schemas.Reimbursement
	err := config.GetDB().Find(&reimbursements).Error
	return reimbursements, err
}

func GetReimbursementById(orderId string) (*hobbies_schemas.Reimbursement, error) {
	var reimbursement hobbies_schemas.Reimbursement
	err := config.GetDB().First(&reimbursement, "order_id = ?", orderId).Error
	if err != nil {
		return nil, err
	}
	return &reimbursement, nil
}

func UpdateReimbursement(reimbursement *hobbies_schemas.Reimbursement) error {
	return config.GetDB().Save(reimbursement).Error
}

func DeleteReimbursement(orderId string) error {
	return config.GetDB().Delete(&hobbies_schemas.Reimbursement{}, "order_id = ?", orderId).Error
}
