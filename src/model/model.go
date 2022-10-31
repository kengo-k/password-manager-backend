package model

import (
	"fmt"
	"time"
)

type Category struct {
	ID        *string    `form:"id" json:"id"`
	Name      *string    `form:"name" json:"name"`
	Desc      *string    `form:"desc" json:"desc"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Password struct {
	ID        *string `form:"id"`
	Category  *string `form:"category"`
	User      *string `form:"user"`
	Password  *string `form:"password"`
	Mail      *string `form:"mail"`
	Note      *string `form:"note"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (p Password) String() string {
	return fmt.Sprintf("{ ID: %s, User: %s, Password: %s, Mail: %s }", *p.ID, *p.User, *p.Password, *p.Mail)
}
