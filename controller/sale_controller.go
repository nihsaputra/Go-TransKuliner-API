package controller

import "github.com/gofiber/fiber/v2"

type SaleController interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
