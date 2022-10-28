package model

import "time"

type Category struct {
	ID   int
	Name string
	Desc string
}

type Password struct {
	ID        int
	Category  int
	User      *string
	Password  *string
	Mail      *string
	Note1     *string
	Note2     *string
	Note3     *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Database struct {
	Categories map[int]Category
	Passwords  map[int]Password
}
