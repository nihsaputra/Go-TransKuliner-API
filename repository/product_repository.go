package repository

import "TransKuliner/model/entity"

type ProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindById(id uint) (entity.Product, error)
	Save(product entity.Product) (entity.Product, error)
	Delete(product entity.Product) error
}
