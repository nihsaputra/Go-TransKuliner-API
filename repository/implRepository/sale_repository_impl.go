package implRepository

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type SaleRepositoryImpl struct {
	Db *gorm.DB
}

func (s *SaleRepositoryImpl) FindAll() ([]entity.Sale, error) {
	var sales []entity.Sale
	tx := s.Db.Preload("Customer").Preload("SaleDetail").Preload("SaleDetail.Product").Preload("SaleDetail.Product.Category").Find(&sales)
	return sales, tx.Error
}

func (s *SaleRepositoryImpl) FindById(id uint) (entity.Sale, error) {
	var sale entity.Sale
	tx := s.Db.Preload("Customer").Preload("SaleDetail").Preload("SaleDetail.Product").Preload("SaleDetail.Product.Category").First(&sale, id)
	return sale, tx.Error
}

func (s *SaleRepositoryImpl) Save(sale entity.Sale) (entity.Sale, error) {
	tx := s.Db.Save(&sale)
	return sale, tx.Error
}

func NewSaleRepository(db *gorm.DB) repository.SaleRepository {
	return &SaleRepositoryImpl{
		Db: db,
	}
}
