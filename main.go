package main

import (
	"TransKuliner/app"
	impl3 "TransKuliner/controller/impl"
	"TransKuliner/repository/impl"
	impl2 "TransKuliner/service/impl"
)

func main() {
	// Database
	mysql := app.NewConnectionMysql()

	// Migration
	app.NewMigration(mysql)

	// Repository
	customerRepository := impl.NewCustomerRepository(mysql)
	categoryRepository := impl.NewCategoryRepository(mysql)
	productRepository := impl.NewProductRepository(mysql)

	// Service
	customerService := impl2.NewCustomerService(customerRepository)
	categoryService := impl2.NewCategoryService(categoryRepository)
	productService := impl2.NewProductService(productRepository, categoryService)

	// Controller
	customerController := impl3.NewCustomerController(customerService)
	categoryController := impl3.NewCategoryController(categoryService)
	productController := impl3.NewProductController(productService)

	// Router
	app.NewCustomerRouter(customerController)
	app.NewCategoryRouter(categoryController)
	app.NewProductRouter(productController)

	router := app.Router
	router.Listen(":8080")

}
