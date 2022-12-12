package model

// Category category table model
type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// CategoryCreateRequest request model for category creation
type CategoryCreateRequest struct {
	Name  *string `form:"name"`
	Order *int    `form:"order"`
}

// CategoryUpdateRequest request model for category update
type CategoryUpdateRequest struct {
	Name  *string `form:"name"`
	Order *int    `form:"order"`
}
