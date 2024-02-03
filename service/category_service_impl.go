package service

import (
	"TransKuliner/handler"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func (c *CategoryServiceImpl) FindAll() []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	findAll, err := c.CategoryRepository.FindAll()
	handler.PanicIfError(err)

	for _, category := range findAll {
		categoryResponse := response.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses
}

func (c *CategoryServiceImpl) FindById(id uint) response.CategoryResponse {
	byId, err := c.CategoryRepository.FindById(id)
	handler.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        byId.ID,
		Name:      byId.Name,
		CreatedAt: byId.CreatedAt,
		UpdatedAt: byId.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Create(request request.CategoryRequest) response.CategoryResponse {
	category := entity.Category{
		Name: request.Name,
	}

	save, err := c.CategoryRepository.Save(category)
	handler.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        save.ID,
		Name:      save.Name,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Update(request request.CategoryUpdateRequest) response.CategoryResponse {
	findById, err := c.CategoryRepository.FindById(request.ID)
	handler.PanicIfError(err)
	findById.Name = request.Name

	save, err := c.CategoryRepository.Save(findById)
	handler.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        save.ID,
		Name:      save.Name,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Delete(id uint) string {
	findById, err := c.CategoryRepository.FindById(id)
	handler.PanicIfError(err)

	err = c.CategoryRepository.Delete(findById)
	handler.PanicIfError(err)

	return "delete successfully"
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
	}
}
