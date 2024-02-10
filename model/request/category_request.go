package request

type CategoryRequest struct {
	Name string `json:"name" validate:"required,min=5"`
}

type CategoryUpdateRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
