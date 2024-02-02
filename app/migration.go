package app

import (
	"TransKuliner/handler"
	"TransKuliner/model/entity"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) {
	err := db.AutoMigrate(entity.Category{}, entity.Product{}, entity.Customer{})
	handler.PanicIfError(err)
}
