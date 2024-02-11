package implService

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/repository"
	"TransKuliner/service"
	"errors"
)

type SaleDetailServiceImpl struct {
	SaleDetailRepository repository.SaleDetailRepository
	SaleService          service.SaleService
	ProductService       service.ProductService
}

func (s *SaleDetailServiceImpl) Create(saleDetailRequest request.SaleDetailRequest) response.SaleDetailResponse {
	productResponse, err := s.ProductService.GetById(saleDetailRequest.ProductId)
	halper.PanicIfError(err)

	if productResponse.Stock < saleDetailRequest.Quantity {
		err := errors.New("stock kurang")
		halper.PanicIfError(err)
	}

	saleDetail := entity.SaleDetail{
		SaleId:    saleDetailRequest.SaleId,
		ProductId: productResponse.ID,
		Quantity:  saleDetailRequest.Quantity,
		Price:     productResponse.Price,
	}
	saveSaleDetail, err := s.SaleDetailRepository.Save(saleDetail)
	halper.PanicIfError(err)

	updateRequest := request.ProductUpdateRequest{
		ID:         productResponse.ID,
		Name:       productResponse.Name,
		Price:      productResponse.Price,
		Stock:      productResponse.Stock - saleDetail.Quantity,
		CategoryID: productResponse.Category.ID,
	}
	s.ProductService.Update(updateRequest)

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
