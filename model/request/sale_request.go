package request

type SaleRequest struct {
	CustomerId uint                      `json:"customer_id"`
	Product    []SaleToSaleDetailRequest `json:"purchase_list"`
}

type SaleToSaleDetailRequest struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
