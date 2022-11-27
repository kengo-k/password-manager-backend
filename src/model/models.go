package model

import (
	"fmt"
	"time"
)

type CategoryUpdateRequest struct {
	Name *string  `form:"name"`
	Order *int `form:"order"`
}

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

func (c Category) String() string {
	return fmt.Sprintf(`{ id: %v, name: %v }`,
		c.ID, c.Name,
	)
}

type PasswordRequest struct {
	ID           *int    `form:"id"`
	CategoryName *string `form:"category_name"`
	User         *string `form:"user"`
	Password     *string `form:"password"`
	Mail         *string `form:"mail"`
	Note         *string `form:"note"`
}

type Password struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Desc      *string  `json:"desc"`
	Category  Category `json:"category"`
	User      *string  `json:"user"`
	Password  *string  `json:"password"`
	Mail      *string  `json:"mail"`
	Note      *string  `json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p Password) String() string {
	return fmt.Sprintf("aaa: %v", p.Category)
	//return fmt.Sprintf(`{ id: %v, name: %v, desc: %v, category: %v, user: %v, password: %v, mail: %v, note: %v, created_at: %v, updated_at: %v }`,
	//	p.ID, p.Name, ifNil(p.Desc), p.Category, ifNil(p.User), ifNil(p.Password), ifNil(p.Mail), ifNil(p.Note), p.CreatedAt, p.UpdatedAt)
}
