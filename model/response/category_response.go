package response

import "time"

type CategoryResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategorySomeResponse struct {
	ID   uint   `json:"category_id"`
	Name string `json:"name"`
}
