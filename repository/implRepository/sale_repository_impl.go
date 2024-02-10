package implRepository

import (
	"TransKuliner/model/entity"
	"TransKuliner/model/request"
	"TransKuliner/repository"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type SaleRepositoryImpl struct {
	Db *gorm.DB
}

func (s *SaleRepositoryImpl) SaveTransaksi(saleRequest request.SaleRequest) error {
	var sale entity.Sale
	var customer entity.Customer

	tx := s.Db.Begin()
	tx.SavePoint("sp1")

	// find CustomerId
	first := tx.First(&customer, saleRequest.CustomerId)
	log.Println(saleRequest.CustomerId)

	if first.Error != nil {
		tx.RollbackTo("sp1")
		return first.Error
	}

	// save Sale
	sale.CustomerId = customer.ID
	sale.CreatedAt = time.Now()
	save := tx.Save(&sale)
	if save.Error != nil {
		tx.RollbackTo("sp1")
		return save.Error
	}

	for _, saleProduct := range saleRequest.Product {
		var product entity.Product

		// find ProductId
		firstDetail := tx.First(&product, saleProduct.ProductId)
		if firstDetail.Error != nil {
			tx.RollbackTo("sp1")
			return firstDetail.Error
		}

		if product.Stock < saleProduct.Quantity {
			tx.RollbackTo("sp1")
			err := errors.New("stock kurang")
			return err
		}

		saleDetail := entity.SaleDetail{
			SaleId:    sale.ID,
			ProductId: product.ID,
			Quantity:  saleProduct.Quantity,
			Price:     product.Price,
		}

		// save SaleDetail
		saveDetail := tx.Save(&saleDetail)
		if saveDetail.Error != nil {
			tx.RollbackTo("sp1")
			return saveDetail.Error
		}

		// update Stock
		product.Stock = product.Stock - saleProduct.Quantity
		updateProduct := tx.Save(&product)
		if updateProduct.Error != nil {
			tx.RollbackTo("sp1")
			return updateProduct.Error
		}
	}

	tx.Commit()
	return nil
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
