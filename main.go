package main

import (
	"TransKuliner/app"
	"TransKuliner/controller/implController"
	"TransKuliner/halper"
	"TransKuliner/repository/implRepository"
	"TransKuliner/routes"
	"TransKuliner/service/implService"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Database
	mysql := app.NewConnectionMysql()

	// Migration
	app.NewMigration(mysql)

	// Repository
	customerRepository := implRepository.NewCustomerRepository(mysql)
	categoryRepository := implRepository.NewCategoryRepository(mysql)
	productRepository := implRepository.NewProductRepository(mysql)
	saleRepository := implRepository.NewSaleRepository(mysql)
	saleDetailRepository := implRepository.NewSaleDetailRepository(mysql)

	// Service
	customerService := implService.NewCustomerService(customerRepository)
	categoryService := implService.NewCategoryService(categoryRepository)
	productService := implService.NewProductService(productRepository, categoryService)
	saleDetailService := implService.NewSaleDetailService(saleDetailRepository, productService)
	saleService := implService.NewSaleService(saleRepository, customerService, saleDetailService)

	// Controller
	customerController := implController.NewCustomerController(customerService)
	categoryController := implController.NewCategoryController(categoryService)
	productController := implController.NewProductController(productService)
	saleController := implController.NewSaleController(saleService)

	router := routes.Routes
	router.Use(recover.New())

	// Routes
	routes.NewCustomerRouter(customerController)
	routes.NewCategoryRouter(categoryController)
	routes.NewProductRouter(productController)
	routes.NewSaleRouter(saleController)
	err := router.Listen(":8080")
	halper.PanicIfError(err)
}
