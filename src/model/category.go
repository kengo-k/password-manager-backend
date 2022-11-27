package model

// category table model
type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// request model for category creation
type CategoryCreateRequest struct {
	Name  *string `form:"name"`
	Order *int    `form:"order"`
}

// request model for category update
type CategoryUpdateRequest struct {
	// TODO passwordに合わせてIDを追加すること
	Name  *string `form:"name"`
	Order *int    `form:"order"`
}
