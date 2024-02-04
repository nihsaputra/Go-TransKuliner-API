package impl

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func (p *ProductRepositoryImpl) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	tx := p.Db.Preload("Category").Find(&products)
	return products, tx.Error
}

func (p *ProductRepositoryImpl) FindById(id uint) (entity.Product, error) {
	var product entity.Product
	tx := p.Db.Preload("Category").First(&product, id)
	return product, tx.Error
}

func (p *ProductRepositoryImpl) Save(product entity.Product) (entity.Product, error) {
	tx := p.Db.Save(&product)
	return product, tx.Error
}

func (p *ProductRepositoryImpl) Delete(product entity.Product) error {
	tx := p.Db.Delete(&product)
	return tx.Error
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &ProductRepositoryImpl{
		Db: db,
	}
}
