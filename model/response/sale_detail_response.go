package response

type SaleDetailResponse struct {
	Product    string `json:"product"`
	Category   string `json:"category"`
	Quantity   uint   `json:"quantity"`
	Price      uint   `json:"price"`
	TotalPrice uint   `json:"total_price"`
}
