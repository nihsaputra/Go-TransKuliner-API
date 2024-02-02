package repository

import "TransKuliner/model/entity"

type CustomerRepository interface {
	FindAll() ([]entity.Customer, error)
	FindById(id uint) (entity.Customer, error)
	Save(customer entity.Customer) (entity.Customer, error)
	Delete(customer entity.Customer) error
}
