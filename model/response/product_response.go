package response

import "time"

type ProductResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Price        uint      `json:"price"`
	Stock        uint      `json:"stock"`
	CategoryName string    `json:"category_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}