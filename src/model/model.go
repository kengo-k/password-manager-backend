package model

import "time"

type Category struct {
	ID   int
	Name string
	Desc string
}

type Password struct {
	ID        int     `form:"id"`
	Category  int     `form:"category"`
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
	Categories map[int]Category
	Passwords  map[int]Password
}
