package request

type ProductRequest struct {
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	Stock      uint   `json:"stock"`
	CategoryID uint   `json:"category_id"`
}

type ProductUpdateRequest struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Price      uint   `json:"price"`
	Stock      uint   `json:"stock"`
	CategoryID uint   `json:"category_id"`
}

type ProductUpdateStockRequest struct {
	ID    uint `json:"id"`
	Stock uint `json:"stock"`
}
