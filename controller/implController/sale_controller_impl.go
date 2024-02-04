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

type SaleControllerImpl struct {
	SaleService service.SaleService
}

func (s *SaleControllerImpl) GetAll(ctx *fiber.Ctx) error {
	saleResponses := s.SaleService.GetAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   saleResponses,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (s *SaleControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	halper.PanicIfError(err)
	saleResponse := s.SaleService.GetById(uint(paramsId))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   saleResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (s *SaleControllerImpl) Create(ctx *fiber.Ctx) error {
	var saleRequest request.SaleRequest

	err := ctx.BodyParser(&saleRequest)
	halper.PanicIfError(err)

	saleResponse := s.SaleService.Create(saleRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATE",
		Data:   saleResponse,
	}
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func NewSaleController(service service.SaleService) controller.SaleController {
	return &SaleControllerImpl{
		SaleService: service,
	}
}
