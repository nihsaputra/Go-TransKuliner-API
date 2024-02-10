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
	customerResponses, err := c.CustomerService.GetAll()
	halper.PanicIfError(err)
	// belum di validasi 404

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponses,
	}
	return ctx.JSON(webResponse)
}

func (c *CustomerControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsId, _ := ctx.ParamsInt("id")
	// belum di validasi 404

	customerResponse, _ := c.CustomerService.GetById(uint(paramsId))
	// belum di validasi 404

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
	// belum di validasi 404

	customerResponse := c.CustomerService.Create(customerRequest)
	// belum di validasi 404

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
	// belum di validasi 404

	customerResponse := c.CustomerService.Update(customerUpdateRequest)
	// belum di validasi 404

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   customerResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CustomerControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsId, _ := ctx.ParamsInt("id")
	// belum di validasi 404

	CustomerResponse := c.CustomerService.Delete(uint(paramsId))
	// belum di validasi 404

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
