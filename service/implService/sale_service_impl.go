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
	SaleRepository    repository.SaleRepository
	ProductService    service.ProductService
	CustomerService   service.CustomerService
	SaleDetailService service.SaleDetailService
}

func (s *SaleServiceImpl) GetAll() []response.SaleResponse {
	sales, errFindAllSale := s.SaleRepository.FindAll()
	halper.PanicIfError(errFindAllSale)

	var saleResponses []response.SaleResponse
	for _, sale := range sales {
		saleResponse := response.SaleResponse{
			ID:                 sale.ID,
			CreatedAt:          sale.CreatedAt,
			CustomerName:       sale.Customer.Name,
			SaleDetailResponse: halper.SaleDetailToSaleDetailResponse(sale.SaleDetail),
		}
		saleResponses = append(saleResponses, saleResponse)
	}

	return saleResponses
}

func (s *SaleServiceImpl) GetById(id uint) response.SaleResponse {
	sale, errFindByIdSale := s.SaleRepository.FindById(id)
	halper.PanicIfError(errFindByIdSale)

	saleResponse := response.SaleResponse{
		ID:                 sale.ID,
		CustomerName:       sale.Customer.Name,
		CreatedAt:          sale.CreatedAt,
		SaleDetailResponse: halper.SaleDetailToSaleDetailResponse(sale.SaleDetail),
	}

	return saleResponse
}

func (s *SaleServiceImpl) Create(saleRequest request.SaleRequest) response.SaleResponse {
	customerResponse, errFindCustomer := s.CustomerService.GetById(saleRequest.CustomerId)
	halper.PanicIfError(errFindCustomer)

	sale := entity.Sale{
		CustomerId: customerResponse.ID,
		CreatedAt:  time.Now(),
	}
	save, errSave := s.SaleRepository.Save(sale)
	halper.PanicIfError(errSave)

	var purchaseList []response.SaleDetailResponse
	for _, saleProduct := range saleRequest.Product {
		saleDetailRequest := request.SaleDetailRequest{
			SaleId:    save.ID,
			ProductId: saleProduct.ProductId,
			Quantity:  saleProduct.Quantity,
		}
		saleDetailCreateResponse := s.SaleDetailService.Create(saleDetailRequest)
		purchaseList = append(purchaseList, saleDetailCreateResponse)
	}

	saleCreateResponse := response.SaleResponse{
		ID:                 save.ID,
		CustomerName:       customerResponse.Name,
		SaleDetailResponse: purchaseList,
	}

	return saleCreateResponse
}

func NewSaleService(repository repository.SaleRepository, customerService service.CustomerService, saleDetailService service.SaleDetailService) service.SaleService {
	return &SaleServiceImpl{
		SaleRepository:    repository,
		CustomerService:   customerService,
		SaleDetailService: saleDetailService,
	}
}
