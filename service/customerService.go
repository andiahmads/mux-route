package service

import (
	"mux-route/domain"
	"mux-route/dto"
	"mux-route/helper"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *helper.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *helper.AppError)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *helper.AppError) {
	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)

	for _, c := range customers {
		response = append(response, c.CustomerFormater())
	}

	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *helper.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.CustomerFormater()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
