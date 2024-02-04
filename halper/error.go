package halper

import (
	"TransKuliner/model/response"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorNotFound(err error, ctx *fiber.Ctx) error {
	webResponse := response.WebResponse{
		Code:   404,
		Status: "Not Found",
		Data:   err.Error(),
	}
	return ctx.Status(http.StatusBadRequest).JSON(webResponse)
}
