package implController

import (
	"TransKuliner/controller"
	"TransKuliner/halper"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type SaleDetailControllerImpl struct {
	SaleDetailService service.SaleDetailService
}

func (s *SaleDetailControllerImpl) GetAll(ctx *fiber.Ctx) error {
	saleDetailResponses := s.SaleDetailService.GetAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   saleDetailResponses,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (s *SaleDetailControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	halper.PanicIfError(err)
	saleDetailResponse := s.SaleDetailService.GetById(uint(paramsId))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   saleDetailResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (s *SaleDetailControllerImpl) Create(ctx *fiber.Ctx) error {
	var saleDetailRequest request.SaleDetailRequest
	err := ctx.BodyParser(&saleDetailRequest)
	halper.PanicIfError(err)

	saleDetailResponse := s.SaleDetailService.Create(saleDetailRequest)

	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATE",
		Data:   saleDetailResponse,
	}

	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func NewSaleDetailController(service service.SaleDetailService) controller.SaleDetailController {
	return &SaleDetailControllerImpl{
		SaleDetailService: service,
	}
}
