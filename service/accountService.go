package service

import (
	"mux-route/domain"
	"mux-route/dto"
	"mux-route/helper"
	"time"
)

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *helper.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *helper.AppError) {
	// validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account := domain.Account{}
	account.AccountId = ""
	account.CustomerId = req.CustomerId
	account.OpeningDate = time.Now().Format("2006-01-02 15:04:05")
	account.AccountType = req.AccountType
	account.Amount = req.Amount
	account.Status = "1"

	res, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := res.AccountResponseDTO()

	return &response, nil
}
