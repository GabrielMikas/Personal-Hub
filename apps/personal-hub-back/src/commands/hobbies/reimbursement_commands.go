package hobbies_commands

import (
	"time"

	hobbies_repositories "github.com/personalhub/personalhub-back/src/repositories/hobbies"
	hobbies_schemas "github.com/personalhub/personalhub-back/src/schemas/hobbies"
)

// CreateReimbursementCommand represents the intent to register a new
// Reimbursement. Unlike Card/Order, a Reimbursement's primary key (OrderId)
// is not generated here — it references an Order that already exists, so
// it must be sent through the request body.
type CreateReimbursementCommand struct {
	OrderId             string    `json:"orderId" binding:"required"`
	StoreName           string    `json:"storeName" binding:"required"`
	StoreContactInfo    string    `json:"storeContactInfo" binding:"required"`
	StoreOrderId        string    `json:"storeOrderId" binding:"required"`
	RequestedAt         time.Time `json:"requestedAt" binding:"required"`
	ReimbursementValue  uint      `json:"reimbursementValue" binding:"required"`
	ReimbursementStatus string    `json:"reimbursementStatus" binding:"required"`
}

func (cmd CreateReimbursementCommand) Execute() (*hobbies_schemas.Reimbursement, error) {
	reimbursement := hobbies_schemas.Reimbursement{
		OrderId:             cmd.OrderId,
		StoreName:           cmd.StoreName,
		StoreContactInfo:    cmd.StoreContactInfo,
		StoreOrderId:        cmd.StoreOrderId,
		RequestedAt:         cmd.RequestedAt,
		ReimbursementValue:  cmd.ReimbursementValue,
		ReimbursementStatus: cmd.ReimbursementStatus,
	}

	if err := hobbies_repositories.CreateReimbursement(&reimbursement); err != nil {
		return nil, err
	}
	return &reimbursement, nil
}

// UpdateReimbursementCommand represents the intent to update an existing
// Reimbursement. OrderId is sent through the request (the URL path) to
// identify which reimbursement to update; only non-nil fields are applied.
type UpdateReimbursementCommand struct {
	OrderId             string     `json:"-"`
	StoreName           *string    `json:"storeName"`
	StoreContactInfo    *string    `json:"storeContactInfo"`
	StoreOrderId        *string    `json:"storeOrderId"`
	RequestedAt         *time.Time `json:"requestedAt"`
	ReimbursementValue  *uint      `json:"reimbursementValue"`
	ReimbursementStatus *string    `json:"reimbursementStatus"`
}

func (cmd UpdateReimbursementCommand) Execute() (*hobbies_schemas.Reimbursement, error) {
	reimbursement, err := hobbies_repositories.GetReimbursementById(cmd.OrderId)
	if err != nil {
		return nil, err
	}

	if cmd.StoreName != nil {
		reimbursement.StoreName = *cmd.StoreName
	}
	if cmd.StoreContactInfo != nil {
		reimbursement.StoreContactInfo = *cmd.StoreContactInfo
	}
	if cmd.StoreOrderId != nil {
		reimbursement.StoreOrderId = *cmd.StoreOrderId
	}
	if cmd.RequestedAt != nil {
		reimbursement.RequestedAt = *cmd.RequestedAt
	}
	if cmd.ReimbursementValue != nil {
		reimbursement.ReimbursementValue = *cmd.ReimbursementValue
	}
	if cmd.ReimbursementStatus != nil {
		reimbursement.ReimbursementStatus = *cmd.ReimbursementStatus
	}

	if err := hobbies_repositories.UpdateReimbursement(reimbursement); err != nil {
		return nil, err
	}
	return reimbursement, nil
}

// DeleteReimbursementCommand represents the intent to delete a
// Reimbursement. OrderId is sent through the request (the URL path).
type DeleteReimbursementCommand struct {
	OrderId string
}

func (cmd DeleteReimbursementCommand) Execute() error {
	if _, err := hobbies_repositories.GetReimbursementById(cmd.OrderId); err != nil {
		return err
	}
	return hobbies_repositories.DeleteReimbursement(cmd.OrderId)
}

// GetReimbursementCommand represents the intent to fetch a single
// Reimbursement. OrderId is sent through the request (the URL path).
type GetReimbursementCommand struct {
	OrderId string
}

func (cmd GetReimbursementCommand) Execute() (*hobbies_schemas.Reimbursement, error) {
	return hobbies_repositories.GetReimbursementById(cmd.OrderId)
}

// ListReimbursementsCommand represents the intent to fetch every
// Reimbursement.
type ListReimbursementsCommand struct{}

func (cmd ListReimbursementsCommand) Execute() ([]hobbies_schemas.Reimbursement, error) {
	return hobbies_repositories.GetAllReimbursements()
}
