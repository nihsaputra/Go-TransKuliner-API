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

func (c *CategoryServiceImpl) GetById(id uint) (response.CategoryResponse, error) {
	byId, err := c.CategoryRepository.FindById(id)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponse{
		ID:        byId.ID,
		Name:      byId.Name,
		CreatedAt: byId.CreatedAt,
		UpdatedAt: byId.UpdatedAt,
	}

	return categoryResponse, nil
}

func (c *CategoryServiceImpl) Create(request request.CategoryRequest) (response.CategoryResponse, error) {
	category := entity.Category{
		Name: request.Name,
	}

	save, err := c.CategoryRepository.Save(category)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponse{
		ID:        save.ID,
		Name:      save.Name,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return categoryResponse, nil
}

func (c *CategoryServiceImpl) Update(request request.CategoryUpdateRequest) (response.CategoryResponse, error) {
	findById, err := c.CategoryRepository.FindById(request.ID)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	findById.Name = request.Name
	save, err := c.CategoryRepository.Save(findById)
	if err != nil {
		return response.CategoryResponse{}, err
	}

	categoryResponse := response.CategoryResponse{
		ID:        save.ID,
		Name:      save.Name,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}
	return categoryResponse, nil
}

func (c *CategoryServiceImpl) Delete(id uint) error {
	findById, err := c.CategoryRepository.FindById(id)
	if err != nil {
		return err
	}

	err = c.CategoryRepository.Delete(findById)
	return err
}

func NewCategoryService(repository repository.CategoryRepository) service.CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
	}
}
