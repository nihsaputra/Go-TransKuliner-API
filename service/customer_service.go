package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type CustomerService interface {
	GetAll() ([]response.CustomerResponse, error)
	GetById(id uint) (response.CustomerResponse, error)
	Create(request request.CustomerRequest) response.CustomerResponse
	Update(request request.CustomerUpdateRequest) response.CustomerResponse
	Delete(id uint) string
}
