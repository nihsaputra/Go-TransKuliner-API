package main

import (
	"TransKuliner/app"
	"TransKuliner/repository"
)

func main() {
	// Database
	mysql := app.NewConnectionMysql()

	// Migration
	app.NewMigration(mysql)

	// Repository
	repository.NewCategoryRepository(mysql)
	repository.NewCustomerRepository(mysql)
	repository.NewProductRepository(mysql)
}
