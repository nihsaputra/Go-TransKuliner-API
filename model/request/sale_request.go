package request

type SaleRequest struct {
	ProductId  uint `json:"product_id"`
	CustomerId uint `json:"customer_id"`
}
