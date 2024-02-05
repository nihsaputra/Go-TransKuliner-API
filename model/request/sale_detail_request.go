package request

type SaleDetailRequest struct {
	SaleId    uint `json:"sale_id"`
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
