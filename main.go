package main

import (
	"TransKuliner/app"
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
	service.NewCustomerService(customerRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	service.NewProductService(productRepository, categoryService)

}
