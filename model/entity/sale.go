package entity

import "time"

type Sale struct {
	ID         uint      `gorm:"column:id;primaryKey"`
	CustomerId uint      `gorm:"column:customer_id"`
	Customer   Customer  `gorm:"foreignKey:customer_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	SaleDetail []SaleDetail
}
