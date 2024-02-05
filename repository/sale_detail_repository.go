package repository

import "TransKuliner/model/entity"

type SaleDetailRepository interface {
	Save(detail entity.SaleDetail) (entity.SaleDetail, error)
}
