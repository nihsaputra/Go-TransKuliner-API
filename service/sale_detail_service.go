package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type SaleDetailService interface {
	GetAll() []response.SaleDetailResponse
	GetById(id uint) response.SaleDetailResponse
	Create(request request.SaleDetailRequest) response.SaleDetailResponse
}
