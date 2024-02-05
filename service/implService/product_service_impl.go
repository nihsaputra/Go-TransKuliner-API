package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	CategoryService   service.CategoryService
}

func (p *ProductServiceImpl) GetAll() []response.ProductResponse {
	var productResponses []response.ProductResponse
	findAll, err := p.ProductRepository.FindAll()
	halper.PanicIfError(err)

	for _, product := range findAll {
		productResponse := response.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			Category:  halper.CategoryToCategorySomeResponse(product.Category),
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}

func (p *ProductServiceImpl) GetById(id uint) response.ProductResponse {
	findById, err := p.ProductRepository.FindById(id)
	halper.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:        findById.ID,
		Name:      findById.Name,
		Price:     findById.Price,
		Stock:     findById.Stock,
		Category:  halper.CategoryToCategorySomeResponse(findById.Category),
		CreatedAt: findById.CreatedAt,
		UpdatedAt: findById.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Create(request request.ProductRequest) response.ProductResponse {
	categoryResponse := p.CategoryService.GetById(request.CategoryID)

	product := entity.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: categoryResponse.ID,
	}

	save, err := p.ProductRepository.Save(product)
	halper.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:        save.ID,
		Name:      save.Name,
		Price:     save.Price,
		Stock:     save.Stock,
		Category:  halper.CategoryToCategorySomeResponse(save.Category),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Update(request request.ProductUpdateRequest) response.ProductResponse {
	findById, err := p.ProductRepository.FindById(request.ID)
	halper.PanicIfError(err)

	categoryResponse := p.CategoryService.GetById(request.CategoryID)

	findById.Name = request.Name
	findById.Price = request.Price
	findById.Stock = request.Stock
	findById.CategoryID = categoryResponse.ID

	save, err := p.ProductRepository.Save(findById)
	halper.PanicIfError(err)

	productResponse := response.ProductResponse{
		ID:        save.ID,
		Name:      save.Name,
		Price:     save.Price,
		Stock:     save.Stock,
		Category:  halper.CategoryToCategorySomeResponse(save.Category),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return productResponse
}

func (p *ProductServiceImpl) Delete(id uint) string {
	findById, err := p.ProductRepository.FindById(id)
	halper.PanicIfError(err)
	err = p.ProductRepository.Delete(findById)
	halper.PanicIfError(err)

	return "delete successfully"
}

func NewProductService(repository repository.ProductRepository, service service.CategoryService) service.ProductService {
	return &ProductServiceImpl{
		ProductRepository: repository,
		CategoryService:   service,
	}
}
