package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type ProductService interface {
	FindAll() []response.ProductResponse
	FindById(id uint) response.ProductResponse
	Create(request request.ProductRequest) response.ProductResponse
	Update(request request.ProductUpdateRequest) response.ProductResponse
	Delete(id uint) string
}
