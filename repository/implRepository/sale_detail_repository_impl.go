package implRepository

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type SaleDetailRepositoryImpl struct {
	Db *gorm.DB
}

func (s *SaleDetailRepositoryImpl) Save(detail entity.SaleDetail) (entity.SaleDetail, error) {
	tx := s.Db.Save(&detail)
	return detail, tx.Error
}

func NewSaleDetailRepository(db *gorm.DB) repository.SaleDetailRepository {
	return &SaleDetailRepositoryImpl{
		Db: db,
	}
}
