package implRepository

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type SaleDetailRepositoryImpl struct {
	Db *gorm.DB
}

func (s *SaleDetailRepositoryImpl) FindAll() ([]entity.SaleDetail, error) {
	var saleDetails []entity.SaleDetail
	tx := s.Db.Find(&saleDetails)
	return saleDetails, tx.Error
}

func (s *SaleDetailRepositoryImpl) FindById(id uint) (entity.SaleDetail, error) {
	var saleDetail entity.SaleDetail
	tx := s.Db.First(&saleDetail, id)
	return saleDetail, tx.Error
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
