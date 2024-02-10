package implRepository

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
	err := c.Db.Find(&categories).Error
	return categories, err
}

func (c *CategoryRepositoryImpl) FindById(id uint) (entity.Category, error) {
	var category entity.Category
	err := c.Db.First(&category, id).Error
	return category, err
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
