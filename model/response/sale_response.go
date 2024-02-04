package response

import "time"

type SaleResponse struct {
	ID        uint                 `json:"id"`
	Product   ProductSomeResponse  `json:"product"`
	Customer  CustomerSomeResponse `json:"customer"`
	CreatedAt time.Time            `json:"created_at"`
}

type SaleSomeResponse struct {
	ID        uint                 `json:"sale_id"`
	Product   ProductSomeResponse  `json:"product"`
	Customer  CustomerSomeResponse `json:"customer"`
	CreatedAt time.Time            `json:"created_at"`
}
