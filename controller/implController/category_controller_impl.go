package implController

import (
	"TransKuliner/controller"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var validate *validator.Validate

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategoryControllerImpl) GetAll(ctx *fiber.Ctx) error {
	categoryResponses := c.CategoryService.GetAll()

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	// validasi paramsId
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	categoryResponse, err := c.CategoryService.GetById(uint(paramsId))
	// validasi findById
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusNotFound).JSON(response.ErrorResponse{
			Code:    404,
			Status:  "NOT_FOUND",
			Message: err.Error(),
		})
	}

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
	// validation bodyParser request
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	// Validation Request
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(categoryRequest)
	if err != nil {
		var messages []string
		for _, e := range err.(validator.ValidationErrors) {
			sprintf := fmt.Sprintf("Error field: %s, on condition: %s", e.Field(), e.ActualTag())
			messages = append(messages, sprintf)
		}
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: messages,
		})
	}

	categoryResponse, err := c.CategoryService.Create(categoryRequest)
	// validasi create
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	var categoryUpdateRequest request.CategoryUpdateRequest
	ctx.BodyParser(&categoryUpdateRequest)

	// Validation Request
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(categoryUpdateRequest)
	if err != nil {
		var messages []string
		for _, e := range err.(validator.ValidationErrors) {
			sprintf := fmt.Sprintf("Error field: %s, on condition: %s", e.Field(), e.ActualTag())
			messages = append(messages, sprintf)
		}
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: messages,
		})
	}

	categoryResponse, err := c.CategoryService.Update(categoryUpdateRequest)
	// validasi update
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	// validasi paramsId
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	err = c.CategoryService.Delete(uint(paramsId))
	// validasi delete
	if err != nil {
		// belum di handle jenis errornya
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "successfully delete categories",
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewCategoryController(service service.CategoryService) controller.CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}
