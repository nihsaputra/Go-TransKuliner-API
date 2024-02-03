package controller

import (
	"TransKuliner/handler"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategoryControllerImpl) GetAll(ctx *fiber.Ctx) error {
	categoryResponses := c.CategoryService.FindAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	if err != nil {
		handler.ErrorNotFound(err, ctx)
	}
	categoryResponse := c.CategoryService.FindById(uint(paramsId))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var categoryRequest request.CategoryRequest
	err := ctx.BodyParser(&categoryRequest)
	handler.PanicIfError(err)

	categoryResponse := c.CategoryService.Create(categoryRequest)
	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	var categoryUpdateRequest request.CategoryUpdateRequest
	err := ctx.BodyParser(&categoryUpdateRequest)
	handler.PanicIfError(err)

	categoryResponse := c.CategoryService.Update(categoryUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	if err != nil {
		handler.ErrorNotFound(err, ctx)
	}
	categoryResponse := c.CategoryService.Delete(uint(paramsId))

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}
