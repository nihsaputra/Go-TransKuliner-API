package response

import "time"

type SaleResponse struct {
	ID        uint                 `json:"id"`
	Product   ProductSomeResponse  `json:"product"`
	Customer  CustomerSomeResponse `json:"customer"`
	CreatedAt time.Time            `json:"created_at"`
}
