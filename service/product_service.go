package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type ProductService interface {
	GetAll() []response.ProductResponse
	GetById(id uint) (response.ProductResponse, error)
	Create(request request.ProductRequest) (response.ProductResponse, error)
	Update(request request.ProductUpdateRequest) (response.ProductResponse, error)
	Delete(id uint) error
}
