package implController

import (
	"TransKuliner/controller"
	"TransKuliner/model/request"
	"TransKuliner/model/response"
	"TransKuliner/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func (c *CustomerControllerImpl) GetAll(ctx *fiber.Ctx) error {
	customerResponses := c.CustomerService.GetAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}
	return ctx.JSON(webResponse)
}

func (c *CustomerControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	// validasi paramsId
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	customerResponse, err := c.CustomerService.GetById(uint(paramsId))
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
		Data:   customerResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CustomerControllerImpl) Create(ctx *fiber.Ctx) error {
	var customerRequest request.CustomerRequest
	err := ctx.BodyParser(&customerRequest)
	// validasi bodyParser
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	customerResponse, err := c.CustomerService.Create(customerRequest)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   customerResponse,
	}
	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (c *CustomerControllerImpl) Update(ctx *fiber.Ctx) error {
	var customerUpdateRequest request.CustomerUpdateRequest
	err := ctx.BodyParser(&customerUpdateRequest)
	// validasi bodyParser
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	customerResponse, err := c.CustomerService.Update(customerUpdateRequest)
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
		Data:   customerResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CustomerControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	// validasi params
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    400,
			Status:  "BAD_REQUEST",
			Message: err.Error(),
		})
	}

	CustomerResponse := c.CustomerService.Delete(uint(paramsId))
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
		Data:   CustomerResponse,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewCustomerController(service service.CustomerService) controller.CustomerController {
	return &CustomerControllerImpl{
		CustomerService: service,
	}
}
