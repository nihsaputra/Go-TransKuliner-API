package main

import (
	"TransKuliner/app"
	"TransKuliner/controller/implController"
	"TransKuliner/repository/implRepository"
	"TransKuliner/service/implService"
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
	implRepository.NewSaleDetailRepository(mysql)

	// Service
	customerService := implService.NewCustomerService(customerRepository)
	categoryService := implService.NewCategoryService(categoryRepository)
	productService := implService.NewProductService(productRepository, categoryService)
	saleService := implService.NewSaleService(saleRepository, productService, customerService)

	// Controller
	customerController := implController.NewCustomerController(customerService)
	categoryController := implController.NewCategoryController(categoryService)
	productController := implController.NewProductController(productService)
	saleController := implController.NewSaleController(saleService)

	// Router
	app.NewCustomerRouter(customerController)
	app.NewCategoryRouter(categoryController)
	app.NewProductRouter(productController)
	app.NewSaleRouter(saleController)

	router := app.Router
	router.Listen(":8080")

}
