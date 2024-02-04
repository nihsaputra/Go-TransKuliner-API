package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
	"time"
)

type SaleServiceImpl struct {
	SaleRepository  repository.SaleRepository
	ProductService  service.ProductService
	CustomerService service.CustomerService
}

func (s *SaleServiceImpl) GetAll() []response.SaleResponse {
	var saleResponses []response.SaleResponse
	findAll, err := s.SaleRepository.FindAll()
	halper.PanicIfError(err)

	for _, sale := range findAll {
		saleResponse := response.SaleResponse{
			ID:        sale.ID,
			Product:   halper.ProductToProductSomeResponse(sale.Product),
			Customer:  halper.CustomerToCustomerSomeResonse(sale.Customer),
			CreatedAt: sale.CreatedAt,
		}
		saleResponses = append(saleResponses, saleResponse)
	}
	return saleResponses
}

func (s *SaleServiceImpl) GetById(id uint) response.SaleResponse {
	findById, err := s.SaleRepository.FindById(id)
	halper.PanicIfError(err)

	saleResponse := response.SaleResponse{
		ID:        findById.ID,
		Product:   halper.ProductToProductSomeResponse(findById.Product),
		Customer:  halper.CustomerToCustomerSomeResonse(findById.Customer),
		CreatedAt: findById.CreatedAt,
	}

	return saleResponse
}

func (s *SaleServiceImpl) Create(request request.SaleRequest) response.SaleResponse {
	productResponse := s.ProductService.FindById(request.ProductId)
	customerResponse, err := s.CustomerService.FindById(request.CustomerId)
	halper.PanicIfError(err)

	sale := entity.Sale{
		ProductId:  productResponse.ID,
		CustomerId: customerResponse.ID,
		CreatedAt:  time.Now(),
	}
	save, errSave := s.SaleRepository.Save(sale)
	halper.PanicIfError(errSave)

	saleResponse := response.SaleResponse{
		ID:        save.ID,
		Product:   halper.ProductToProductSomeResponse(save.Product),
		Customer:  halper.CustomerToCustomerSomeResonse(save.Customer),
		CreatedAt: save.CreatedAt,
	}

	return saleResponse
}

func NewSaleService(repository repository.SaleRepository, productService service.ProductService, customerService service.CustomerService) service.SaleService {
	return &SaleServiceImpl{
		SaleRepository:  repository,
		ProductService:  productService,
		CustomerService: customerService,
	}
}
