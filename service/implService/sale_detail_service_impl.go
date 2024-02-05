package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
)

type SaleDetailServiceImpl struct {
	SaleDetailRepository repository.SaleDetailRepository
	SaleService          service.SaleService
	ProductService       service.ProductService
}

func (s *SaleDetailServiceImpl) Create(request request.SaleDetailRequest) response.SaleDetailResponse {
	productResponse := s.ProductService.GetById(request.ProductId)

	saleDetail := entity.SaleDetail{
		SaleId:    request.SaleId,
		ProductId: productResponse.ID,
		Quantity:  request.Quantity,
		Price:     productResponse.Price,
	}
	saveSaleDetail, err := s.SaleDetailRepository.Save(saleDetail)
	halper.PanicIfError(err)

	saleDetailCreateResponse := response.SaleDetailResponse{
		Product:    productResponse.Name,
		Category:   productResponse.Category.Name,
		Price:      saveSaleDetail.Price,
		Quantity:   saveSaleDetail.Quantity,
		TotalPrice: saveSaleDetail.Price * saveSaleDetail.Quantity,
	}
	return saleDetailCreateResponse
}

func NewSaleDetailService(repository repository.SaleDetailRepository, productService service.ProductService) service.SaleDetailService {
	return &SaleDetailServiceImpl{
		SaleDetailRepository: repository,
		ProductService:       productService,
	}
}
