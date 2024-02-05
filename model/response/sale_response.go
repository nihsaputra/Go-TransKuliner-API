package response

import "time"

type SaleResponse struct {
	ID                 uint                 `json:"id"`
	CustomerName       string               `json:"customer"`
	CreatedAt          time.Time            `json:"created_at"`
	SaleDetailResponse []SaleDetailResponse `json:"purchase_list"`
}
