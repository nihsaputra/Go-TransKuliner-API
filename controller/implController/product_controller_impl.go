package implController

import (
	"TransKuliner/controller"
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
	// validasi paramsId
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	productResponse, err := p.ProductService.GetById(uint(paramsId))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

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
	// validasi bodyParser
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	productResponse, err := p.ProductService.Create(productRequest)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

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
	// validasi bodyParser
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	productResponse, err := p.ProductService.Update(productUpdateRequest)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (p *ProductControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	// validasi paramsId
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	err = p.ProductService.Delete(uint(paramsId))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "successfully delete",
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewProductController(service service.ProductService) controller.ProductController {
	return &ProductControllerImpl{
		ProductService: service,
	}
}
