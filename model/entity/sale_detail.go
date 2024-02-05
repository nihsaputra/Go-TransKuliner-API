package entity

type SaleDetail struct {
	ID        uint    `gorm:"column:id;primaryKey"`
	SaleId    uint    `gorm:"column:sale_id"`
	Sale      Sale    `gorm:"column:sale_id"`
	ProductId uint    `gorm:"column:product_id"`
	Product   Product `gorm:"foreignKey:product_id"`
	Quantity  uint    `gorm:"column:quantity"`
	Price     uint    `gorm:"column:price"`
}
