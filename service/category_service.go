package service

import (
	"TransKuliner/model/request"
	"TransKuliner/model/response"
)

type CategoryService interface {
	GetAll() []response.CategoryResponse
	GetById(id uint) response.CategoryResponse
	Create(request request.CategoryRequest) response.CategoryResponse
	Update(request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(id uint) string
}
