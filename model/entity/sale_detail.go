package entity

type SaleDetail struct {
	ID       uint `gorm:"column:id;primaryKey"`
	SaleId   uint `gorm:"column:sale_id"`
	Sale     Sale `gorm:"column:sale_id"`
	Quantity uint `gorm:"column:quantity"`
	Price    uint `gorm:"column:price"`
}
