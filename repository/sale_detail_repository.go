package repository

import "TransKuliner/model/entity"

type SaleDetailRepository interface {
	FindAll() ([]entity.SaleDetail, error)
	FindById(id uint) (entity.SaleDetail, error)
	Save(detail entity.SaleDetail) (entity.SaleDetail, error)
}
