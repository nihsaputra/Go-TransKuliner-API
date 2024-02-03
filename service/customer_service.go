package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type CustomerService interface {
	FindAll() []response.CustomerResponse
	FindById(id uint) response.CustomerResponse
	Create(request request.CustomerRequest) response.CustomerResponse
	Update(request request.CustomerUpdateRequest) response.CustomerResponse
	Delete(id uint) string
}
