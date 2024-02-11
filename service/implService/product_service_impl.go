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

func (p *ProductServiceImpl) GetById(id uint) (response.ProductResponse, error) {
	findById, err := p.ProductRepository.FindById(id)
	if err != nil {
		return response.ProductResponse{}, err
	}

	productResponse := response.ProductResponse{
		ID:        findById.ID,
		Name:      findById.Name,
		Price:     findById.Price,
		Stock:     findById.Stock,
		Category:  halper.CategoryToCategorySomeResponse(findById.Category),
		CreatedAt: findById.CreatedAt,
		UpdatedAt: findById.UpdatedAt,
	}

	return productResponse, nil
}

func (p *ProductServiceImpl) Create(request request.ProductRequest) (response.ProductResponse, error) {
	categoryResponse, err := p.CategoryService.GetById(request.CategoryID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	product := entity.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: categoryResponse.ID,
	}

	save, err := p.ProductRepository.Save(product)
	if err != nil {
		return response.ProductResponse{}, err
	}

	productResponse := response.ProductResponse{
		ID:        save.ID,
		Name:      save.Name,
		Price:     save.Price,
		Stock:     save.Stock,
		Category:  halper.CategoryToCategorySomeResponse(save.Category),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return productResponse, nil
}

func (p *ProductServiceImpl) Update(request request.ProductUpdateRequest) (response.ProductResponse, error) {
	findById, err := p.ProductRepository.FindById(request.ID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	categoryResponse, err := p.CategoryService.GetById(request.CategoryID)
	if err != nil {
		return response.ProductResponse{}, err
	}

	findById.Name = request.Name
	findById.Price = request.Price
	findById.Stock = request.Stock
	findById.CategoryID = categoryResponse.ID

	save, err := p.ProductRepository.Save(findById)
	if err != nil {
		return response.ProductResponse{}, err
	}

	productResponse := response.ProductResponse{
		ID:        save.ID,
		Name:      save.Name,
		Price:     save.Price,
		Stock:     save.Stock,
		Category:  halper.CategoryToCategorySomeResponse(save.Category),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}

	return productResponse, nil
}

func (p *ProductServiceImpl) Delete(id uint) error {
	findById, err := p.ProductRepository.FindById(id)
	if err != nil {
		return err
	}

	err = p.ProductRepository.Delete(findById)
	return err
}

func NewProductService(repository repository.ProductRepository, service service.CategoryService) service.ProductService {
	return &ProductServiceImpl{
		ProductRepository: repository,
		CategoryService:   service,
	}
}
