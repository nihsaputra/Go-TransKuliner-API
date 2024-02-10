package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type CategoryService interface {
	GetAll() []response.CategoryResponse
	GetById(id uint) (response.CategoryResponse, error)
	Create(request request.CategoryRequest) (response.CategoryResponse, error)
	Update(request request.CategoryUpdateRequest) (response.CategoryResponse, error)
	Delete(id uint) error
}
