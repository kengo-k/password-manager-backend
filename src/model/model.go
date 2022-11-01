package model

import (
	"fmt"
	"time"
)

type CategoryRequest struct {
	Name string  `form:"name"`
	Desc *string `form:"desc"`
}

type Category struct {
	Name      string    `json:"name"`
	Desc      *string   `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c Category) String() string {
	return fmt.Sprintf(`{ name: %v, desc: %v, created_at: %v, updated_at: %v }`,
		ifNil(&c.Name), ifNil(c.Desc), c.CreatedAt, c.UpdatedAt)
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

func ifNil(ps *string) string {
	if ps == nil {
		return "<nil>"
	}
	return fmt.Sprintf("\"%s\"", *ps)
}

func (p Password) String() string {
	return fmt.Sprintf("aaa: %v", p.Category)
	//return fmt.Sprintf(`{ id: %v, name: %v, desc: %v, category: %v, user: %v, password: %v, mail: %v, note: %v, created_at: %v, updated_at: %v }`,
	//	p.ID, p.Name, ifNil(p.Desc), p.Category, ifNil(p.User), ifNil(p.Password), ifNil(p.Mail), ifNil(p.Note), p.CreatedAt, p.UpdatedAt)
}
