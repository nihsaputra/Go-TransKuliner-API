package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type CustomerService interface {
	GetAll() []response.CustomerResponse
	GetById(id uint) (response.CustomerResponse, error)
	Create(request request.CustomerRequest) (response.CustomerResponse, error)
	Update(request request.CustomerUpdateRequest) (response.CustomerResponse, error)
	Delete(id uint) error
}
