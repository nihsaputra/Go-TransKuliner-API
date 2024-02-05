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

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func (p *ProductControllerImpl) GetAll(ctx *fiber.Ctx) error {
	productResponses := p.ProductService.GetAll()

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponses,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (p *ProductControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	if err != nil {
		return halper.ErrorNotFound(err, ctx)
	}

	productResponse := p.ProductService.GetById(uint(paramsId))

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (p *ProductControllerImpl) Create(ctx *fiber.Ctx) error {
	var productRequest request.ProductRequest
	err := ctx.BodyParser(&productRequest)
	halper.PanicIfError(err)

	productResponse := p.ProductService.Create(productRequest)

	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATE",
		Data:   productResponse,
	}

	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (p *ProductControllerImpl) Update(ctx *fiber.Ctx) error {
	var productUpdateRequest request.ProductUpdateRequest
	err := ctx.BodyParser(&productUpdateRequest)
	halper.PanicIfError(err)

	productResponse := p.ProductService.Update(productUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (p *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	if err != nil {
		halper.ErrorNotFound(err, ctx)
	}

	productResponse := p.ProductService.Delete(uint(paramsId))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewProductController(service service.ProductService) controller.ProductController {
	return &ProductControllerImpl{
		ProductService: service,
	}
}
