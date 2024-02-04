package app

import (
	"TransKuliner/halper"
	"TransKuliner/model/entity"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		entity.Category{}, entity.Product{}, entity.Customer{},
		entity.Sale{}, entity.SaleDetail{},
	)

	halper.PanicIfError(err)
}
