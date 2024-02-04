package impl

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func (c *CategoryRepositoryImpl) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	tx := c.Db.Find(&categories)
	return categories, tx.Error
}

func (c *CategoryRepositoryImpl) FindById(id uint) (entity.Category, error) {
	var category entity.Category
	tx := c.Db.First(&category, id)
	return category, tx.Error
}

func (c *CategoryRepositoryImpl) Save(category entity.Category) (entity.Category, error) {
	tx := c.Db.Save(&category)
	return category, tx.Error
}

func (c *CategoryRepositoryImpl) Delete(category entity.Category) error {
	tx := c.Db.Delete(&category)
	return tx.Error
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepositoryImpl{
		Db: db,
	}
}
