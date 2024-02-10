package routes

import (
	"TransKuliner/controller"
	"github.com/gofiber/fiber/v2"
)

var Routes = *fiber.New()

func NewCustomerRouter(customerController controller.CustomerController) {
	Routes.Get("/customers", customerController.GetAll)
	Routes.Get("/customers/:id", customerController.GetById)
	Routes.Post("/customers", customerController.Create)
	Routes.Put("/customers", customerController.Update)
	Routes.Delete("/customers/:id", customerController.Delete)
}

func NewCategoryRouter(categoryController controller.CategoryController) {
	Routes.Get("/categories", categoryController.GetAll)
	Routes.Get("/categories/:id", categoryController.GetById)
	Routes.Post("/categories", categoryController.Create)
	Routes.Put("/categories", categoryController.Update)
	Routes.Delete("/categories/:id", categoryController.Delete)
}

func NewProductRouter(productController controller.ProductController) {
	Routes.Get("/products", productController.GetAll)
	Routes.Get("/products/:id", productController.GetById)
	Routes.Post("/products", productController.Create)
	Routes.Put("/products", productController.Update)
	Routes.Delete("/products/:id", productController.Delete)
}

func NewSaleRouter(saleController controller.SaleController) {
	Routes.Get("/sales", saleController.GetAll)
	Routes.Get("/sales/:id", saleController.GetById)
	Routes.Post("/sales", saleController.Create)
}
