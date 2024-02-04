package response

import "time"

type CustomerResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomerSomeResponse struct {
	ID   uint   `json:"customer_id"`
	Name string `json:"name"`
}
