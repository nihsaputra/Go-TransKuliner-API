package impl

import (
	"TransKuliner/model/entity"
	"TransKuliner/repository"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

func (c *CustomerRepositoryImpl) FindAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	tx := c.Db.Find(&customers)
	return customers, tx.Error
}

func (c *CustomerRepositoryImpl) FindById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	tx := c.Db.First(&customer, id)
	return customer, tx.Error
}

func (c *CustomerRepositoryImpl) Save(customer entity.Customer) (entity.Customer, error) {
	tx := c.Db.Save(&customer)
	return customer, tx.Error
}

func (c *CustomerRepositoryImpl) Delete(customer entity.Customer) error {
	tx := c.Db.Delete(&customer)
	return tx.Error
}

func NewCustomerRepository(db *gorm.DB) repository.CustomerRepository {
	return &CustomerRepositoryImpl{
		Db: db,
	}
}
