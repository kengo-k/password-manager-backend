package model

import (
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
	Note1     *string `form:"note1"`
	Note2     *string `form:"note2"`
	Note3     *string `form:"note3"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Database struct {
	Categories map[string]*Category
	Passwords  map[string]*Password
}