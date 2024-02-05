package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type SaleDetailService interface {
	Create(request request.SaleDetailRequest) response.SaleDetailResponse
}
