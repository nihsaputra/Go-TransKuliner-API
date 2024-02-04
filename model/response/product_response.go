package response

import "time"

type ProductResponse struct {
	ID        uint                 `json:"id"`
	Name      string               `json:"name"`
	Price     uint                 `json:"price"`
	Stock     uint                 `json:"stock"`
	Category  CategorySomeResponse `json:"category"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
type ProductSomeResponse struct {
	ID       uint                 `json:"product_id"`
	Name     string               `json:"name"`
	Price    uint                 `json:"price"`
	Stock    uint                 `json:"stock"`
	Category CategorySomeResponse `json:"category"`
}
