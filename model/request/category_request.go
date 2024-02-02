package request

type CategoryRequest struct {
	Name string `json:"name"`
}

type CategoryUpdateRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
