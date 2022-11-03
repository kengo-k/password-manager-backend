package model

import (
	"fmt"
	"strings"
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

type Database struct {
	Passwords  map[int]*Password
	Categories map[string]*Category
}

func NewDatabase() *Database {
	return &Database{
		Passwords:  map[int]*Password{},
		Categories: map[string]*Category{},
	}
}

// 旧形式パスワードを読み込む
// ※移行が完了したら不要になる
func (d *Database) Init(mdLines []string) error {

	foundCategory := false
	foundHeader := false

	splitColumns := func(line string) []string {
		ret := []string{}
		for _, column := range strings.Split(line, "|") {
			column = strings.TrimSpace(column)
			if len(column) > 0 {
				ret = append(ret, column)
			}
		}
		return ret
	}
	var c Category
	pid := 0
	for _, l := range mdLines {
		// 空行の場合はスキップする
		if len(l) == 0 {
			continue
		}
		// #で始まるコメント行の場合はカテゴリ名が記載されている
		if strings.HasPrefix(l, "#") {
			foundCategory = true
			_, categoryName, ok := strings.Cut(l, "#")
			if !ok {
				return fmt.Errorf("failed to get catgory name")
			}
			c = Category{
				Name: categoryName,
				Desc: nil,
			}
			d.Categories[categoryName] = &c
			continue
		}
		if foundCategory {
			foundCategory = false
			foundHeader = true
			continue
		}
		if foundHeader {
			foundHeader = false
			continue
		}

		pid++
		columns := splitColumns(l)
		if len(columns) != 5 && len(columns) != 6 {
			return fmt.Errorf("faild to load, invalid column length: %v", len(columns))
		}

		if len(columns) == 5 {
			p := &Password{
				ID:       pid,
				Name:     columns[0],
				Desc:     &columns[1],
				Category: c,
				User:     &columns[2],
				Password: &columns[3],
				Mail:     nil,
				Note:     &columns[4],
			}
			d.Passwords[p.ID] = p
		}
		if len(columns) == 6 {
			p := &Password{
				ID:       pid,
				Name:     columns[0],
				Desc:     &columns[1],
				Category: c,
				User:     &columns[2],
				Password: &columns[3],
				Mail:     &columns[4],
				Note:     &columns[5],
			}
			d.Passwords[p.ID] = p
		}
	}

	return nil
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
