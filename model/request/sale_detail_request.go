package request

type SaleDetailRequest struct {
	SaleId   uint `json:"sale_id"`
	Quantity uint `json:"quantity"`
	Price    uint `json:"price"`
}
