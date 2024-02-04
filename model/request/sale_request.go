package request

import "time"

type SaleRequest struct {
	ProductId  uint      `json:"product_id"`
	CustomerId uint      `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}
