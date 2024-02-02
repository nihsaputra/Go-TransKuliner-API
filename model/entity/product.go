package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string   `gorm:"column:name"`
	Price      uint     `gorm:"column:price"`
	Stock      uint     `gorm:"column:stock"`
	CategoryID uint     `gorm:"column:category_id"`
	Category   Category `gorm:"foreignKey:category_id"`
}
