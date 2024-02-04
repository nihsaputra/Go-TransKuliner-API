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

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func (c *CustomerControllerImpl) GetAll(ctx *fiber.Ctx) error {
	customerResponses, err := c.CustomerService.FindAll()

	if err != nil {
		return halper.ErrorNotFound(err, ctx)
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}
	return ctx.JSON(webResponse)
}

func (c *CustomerControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")

	if err != nil {
		return halper.ErrorNotFound(err, ctx)
	}

	customerResponse, errFind := c.CustomerService.FindById(uint(paramsId))
	if errFind != nil {
		return halper.ErrorNotFound(errFind, ctx)
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
	halper.PanicIfError(err)

	customerResponse := c.CustomerService.Create(customerRequest)

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
	halper.PanicIfError(err)

	customerResponse := c.CustomerService.Update(customerUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CustomerControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, err := ctx.ParamsInt("id")
	if err != nil {
		halper.ErrorNotFound(err, ctx)
	}

	CustomerResponse := c.CustomerService.Delete(uint(paramsId))
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
