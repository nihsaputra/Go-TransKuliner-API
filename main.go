package main

import (
	"TransKuliner/app"
	"TransKuliner/controller"
	"TransKuliner/repository"
	"TransKuliner/service"
)

func main() {
	// Database
	mysql := app.NewConnectionMysql()

	// Migration
	app.NewMigration(mysql)

	// Repository
	customerRepository := repository.NewCustomerRepository(mysql)
	categoryRepository := repository.NewCategoryRepository(mysql)
	productRepository := repository.NewProductRepository(mysql)

	// Service
	customerService := service.NewCustomerService(customerRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository, categoryService)

	// Controller
	customerController := controller.NewCustomerController(customerService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	// Router
	app.NewCustomerRouter(customerController)
	app.NewCategoryRouter(categoryController)
	app.NewProductRouter(productController)

	router := app.Router
	router.Listen(":8080")

}
