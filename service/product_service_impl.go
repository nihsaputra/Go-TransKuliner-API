package service

import (
	"TransKuliner/handler"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	CategoryService   CategoryService
}

func (p *ProductServiceImpl) FindAll() []response.ProductResponse {
	var productResponses []response.ProductResponse
	findAll, err := p.ProductRepository.FindAll()
	handler.PanicIfError(err)

	for _, product := range findAll {
		productResponse := response.ProductResponse{
			ID:           product.ID,
			Name:         product.Name,
			Price:        product.Price,
			Stock:        product.Stock,
			CategoryName: product.Category.Name,
			CreatedAt:    product.CreatedAt,
			UpdatedAt:    product.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}

func (p *ProductServiceImpl) FindById(id uint) response.ProductResponse {
	findById, err := p.ProductRepository.FindById(id)
	handler.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:           findById.ID,
		Name:         findById.Name,
		Price:        findById.Price,
		Stock:        findById.Stock,
		CategoryName: findById.Category.Name,
		CreatedAt:    findById.CreatedAt,
		UpdatedAt:    findById.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Create(request request.ProductRequest) response.ProductResponse {
	categoryResponse := p.CategoryService.FindById(request.CategoryID)

	product := entity.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: categoryResponse.ID,
	}

	save, err := p.ProductRepository.Save(product)
	handler.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:           save.ID,
		Name:         save.Name,
		Price:        save.Price,
		Stock:        save.Stock,
		CategoryName: categoryResponse.Name,
		CreatedAt:    save.CreatedAt,
		UpdatedAt:    save.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Update(request request.ProductUpdateRequest) response.ProductResponse {
	findById, err := p.ProductRepository.FindById(request.ID)
	handler.PanicIfError(err)

	categoryResponse := p.CategoryService.FindById(request.CategoryID)

	findById.Name = request.Name
	findById.Price = request.Price
	findById.Stock = request.Stock
	findById.CategoryID = categoryResponse.ID

	save, err := p.ProductRepository.Save(findById)
	handler.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:           save.ID,
		Name:         save.Name,
		Price:        save.Price,
		Stock:        save.Stock,
		CategoryName: categoryResponse.Name,
		CreatedAt:    save.CreatedAt,
		UpdatedAt:    save.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Delete(id uint) string {
	findById, err := p.ProductRepository.FindById(id)
	handler.PanicIfError(err)
	err = p.ProductRepository.Delete(findById)
	handler.PanicIfError(err)

	return "delete successfully"
}

func NewProductService(repository repository.ProductRepository, service CategoryService) ProductService {
	return &ProductServiceImpl{
		ProductRepository: repository,
		CategoryService:   service,
	}
}
