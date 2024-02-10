package repository

import (
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
)

type SaleRepository interface {
	FindAll() ([]entity.Sale, error)
	FindById(id uint) (entity.Sale, error)
	Save(sale entity.Sale) (entity.Sale, error)
	SaveTransaksi(saleRequst request.SaleRequest) error
}
