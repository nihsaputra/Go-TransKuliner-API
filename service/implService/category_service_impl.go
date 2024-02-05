package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func (c *CategoryServiceImpl) GetAll() []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	findAll, err := c.CategoryRepository.FindAll()
	halper.PanicIfError(err)

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

func (c *CategoryServiceImpl) GetById(id uint) response.CategoryResponse {
	byId, err := c.CategoryRepository.FindById(id)
	halper.PanicIfError(err)

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
	halper.PanicIfError(err)

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
	halper.PanicIfError(err)
	findById.Name = request.Name

	save, err := c.CategoryRepository.Save(findById)
	halper.PanicIfError(err)

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
	halper.PanicIfError(err)

	err = c.CategoryRepository.Delete(findById)
	halper.PanicIfError(err)

	return "delete successfully"
}

func NewCategoryService(repository repository.CategoryRepository) service.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
	}
}
