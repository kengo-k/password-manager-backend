package model

import (
	"fmt"
	"sort"
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

func (d *Database) Serialize() [][]*Password {
	ret := [][]*Password{}
	cmap := map[string]*[]*Password{}
	for cname := range d.Categories {
		passwords := []*Password{}
		ret = append(ret, passwords)
		cmap[cname] = &passwords
	}
	for _, p := range d.Passwords {
		passwords := cmap[p.Category.Name]
		*passwords = append(*passwords, p)
		cmap[p.Category.Name] = passwords
	}
	sort.Slice(ret, func(a int, b int) bool {
		p1 := ret[a]
		p2 := ret[b]
		return p1[0].Category.Name < p2[0].Category.Name
	})
	fmt.Printf("ret len: %v", len(ret))
	return ret
}

func (d *Database) ConvertMarkdown() []string {
	// markdown := []string{}
	// for categoryName, categoryPasswords := range passwords {
	// 	head := categoryPasswords[0]
	// 	categoryLine := fmt.Sprintf("# %s: %s", categoryName, *head.Category.Desc)
	// 	headerLine := "|id|name|desc|user|password|mail|note|"
	// 	separatorLine := "|---|---|---|---|---|---|---|"
	// 	markdown = append(markdown, categoryLine)
	// 	markdown = append(markdown, headerLine)
	// 	markdown = append(markdown, separatorLine)
	// 	for _, pwd := range categoryPasswords {
	// 		line := fmt.Sprintf("|%v|%v|%v|%v|%v|%v|%v|",
	// 			pwd.ID, pwd.Name, *pwd.Desc, *pwd.User, *pwd.Password, *pwd.Mail, *pwd.Note)
	// 		markdown = append(markdown, line)
	// 	}
	// }
	// return markdown
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
