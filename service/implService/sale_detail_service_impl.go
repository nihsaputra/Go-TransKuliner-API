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
}

func (s *SaleDetailServiceImpl) GetAll() []response.SaleDetailResponse {
	var saleDetailResponses []response.SaleDetailResponse
	saleDetails, err := s.SaleDetailRepository.FindAll()
	halper.PanicIfError(err)

	for _, saleDetail := range saleDetails {
		saleDetailResponse := response.SaleDetailResponse{
			ID:         saleDetail.ID,
			Sale:       halper.SaleToSaleResponse(saleDetail.Sale),
			Quantity:   saleDetail.Quantity,
			Price:      saleDetail.Price,
			TotalPrice: saleDetail.Quantity * saleDetail.Price,
		}
		saleDetailResponses = append(saleDetailResponses, saleDetailResponse)
	}
	return saleDetailResponses
}

func (s *SaleDetailServiceImpl) GetById(id uint) response.SaleDetailResponse {
	saleDetail, err := s.SaleDetailRepository.FindById(id)
	halper.PanicIfError(err)

	detailResponse := response.SaleDetailResponse{
		ID:         saleDetail.ID,
		Sale:       halper.SaleToSaleResponse(saleDetail.Sale),
		Quantity:   saleDetail.Quantity,
		Price:      saleDetail.Price,
		TotalPrice: saleDetail.Quantity * saleDetail.Price,
	}

	return detailResponse
}

func (s *SaleDetailServiceImpl) Create(request request.SaleDetailRequest) response.SaleDetailResponse {
	saleResponse := s.SaleService.GetById(request.SaleId)

	saleDetail := entity.SaleDetail{
		SaleId:   saleResponse.ID,
		Quantity: request.Quantity,
		Price:    request.Price,
	}
	saveSaleDetail, err := s.SaleDetailRepository.Save(saleDetail)
	halper.PanicIfError(err)

	saleDetailResponse := response.SaleDetailResponse{
		ID:         saveSaleDetail.ID,
		Sale:       saleResponse,
		Price:      saveSaleDetail.Price,
		Quantity:   saveSaleDetail.Quantity,
		TotalPrice: saveSaleDetail.Price * saveSaleDetail.Quantity,
	}

	return saleDetailResponse
}

func NewSaleDetailService(repository repository.SaleDetailRepository, saleService service.SaleService) service.SaleDetailService {
	return &SaleDetailServiceImpl{
		SaleDetailRepository: repository,
		SaleService:          saleService,
	}
}
