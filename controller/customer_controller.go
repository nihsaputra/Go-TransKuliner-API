package controller

import "github.com/gofiber/fiber/v2"

type CustomerController interface {
	GetAll(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
