package model

import "time"

// password table model
type Password struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      *string   `json:"desc"`
	Category  *Category `json:"category"`
	User      *string   `json:"user"`
	Password  *string   `json:"password"`
	Mail      *string   `json:"mail"`
	Note      *string   `json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// request model for password creation
type PasswordCreateRequest struct {
	Name       *string `form:"name"`
	Desc       *string `form:"desc"`
	CategoryID *string `form:"category_id"`
	User       *string `form:"user"`
	Password   *string `form:"password"`
	Mail       *string `form:"mail"`
	Note       *string `form:"note"`
}

// request model for password update
type PasswordUpdateRequest struct {
	ID         int     `form:"id"`
	Name       *string `form:"name"`
	Desc       *string `form:"desc"`
	CategoryID *string `form:"category_id"`
	User       *string `form:"user"`
	Password   *string `form:"password"`
	Mail       *string `form:"mail"`
	Note       *string `form:"note"`
}
