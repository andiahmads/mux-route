package dto

import (
	"mux-route/helper"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type NewAccountResponse struct {
	AccountId string `json:"account_id"`
}

func (r NewAccountRequest) Validate() *helper.AppError {
	if r.Amount < 5000 {
		return helper.NewValidationError("To open new account need deposit atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && r.AccountType != "checking" {
		return helper.NewValidationError("account type should be checking or saving")
	}
	return nil
}
