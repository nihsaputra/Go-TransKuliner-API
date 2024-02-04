package app

import (
	"TransKuliner/controller"
	"github.com/gofiber/fiber/v2"
)

var Router = *fiber.New()

func NewCustomerRouter(customerController controller.CustomerController) {
	Router.Get("/customers", customerController.GetAll)
	Router.Get("/customers/:id", customerController.GetById)
	Router.Post("/customers", customerController.Create)
	Router.Put("/customers", customerController.Update)
	Router.Delete("/customers/:id", customerController.Delete)
}

func NewCategoryRouter(categoryController controller.CategoryController) {
	Router.Get("/categories", categoryController.GetAll)
	Router.Get("/categories/:id", categoryController.GetById)
	Router.Post("/categories", categoryController.Create)
	Router.Put("/categories", categoryController.Update)
	Router.Delete("/categories/:id", categoryController.Delete)
}

func NewProductRouter(productController controller.ProductController) {
	Router.Get("/products", productController.GetAll)
	Router.Get("/products/:id", productController.GetById)
	Router.Post("/products", productController.Create)
	Router.Put("/products", productController.Update)
	Router.Delete("/products/:id", productController.Delete)
}

func NewSaleRouter(saleController controller.SaleController) {
	Router.Get("/sales", saleController.GetAll)
	Router.Get("/sales/:id", saleController.GetById)
	Router.Post("/sales", saleController.Create)
}

func NewSaleDetailRouter(saleDetailController controller.SaleDetailController) {
	Router.Get("/sale-details", saleDetailController.GetAll)
	Router.Get("/sale-details/:id", saleDetailController.GetById)
	Router.Post("/sale-details", saleDetailController.Create)
}
