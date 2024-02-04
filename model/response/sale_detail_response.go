package response

type SaleDetailResponse struct {
	ID         uint         `json:"id"`
	Sale       SaleResponse `json:"sale"`
	Quantity   uint         `json:"quantity"`
	Price      uint         `json:"price"`
	TotalPrice uint         `json:"total_price"`
}
