package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type SaleService interface {
	GetAll() []response.SaleResponse
	GetById(id uint) response.SaleResponse
	Create(request request.SaleRequest) response.SaleResponse
}
