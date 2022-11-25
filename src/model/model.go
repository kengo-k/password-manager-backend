package model

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CategoryRequest struct {
	Name string  `form:"name"`
	Desc *string `form:"desc"`
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
			_, categoryLine, ok := strings.Cut(l, "#")
			if !ok {
				return fmt.Errorf("failed to get catgory line")
			}
			categoryLine = strings.TrimSpace(categoryLine)
			categoryId, attrLine, ok := strings.Cut(categoryLine, ":")
			if !ok {
				return fmt.Errorf("failed to get categoryId")
			}
			attrLine = strings.TrimSpace(attrLine)
			attrs := strings.Split(attrLine, ",")
			attrMap := map[string]string{}
			for _, attr := range attrs {
				k, v, ok := strings.Cut(attr, "=")
				if !ok {
					return fmt.Errorf("failed to parse category attribute")
				}
				attrMap[k] = v
			}
			orderVal, err := strconv.ParseInt(attrMap["order"], 10, 32)
			if err != nil {
				return fmt.Errorf("failed to parse category order to number")
			}
			c = Category{
				ID:    categoryId,
				Name:  attrMap["name"],
				Order: int(orderVal),
			}
			c2 := c
			d.Categories[categoryId] = &c2
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
	return serialize(d.Categories, d.Passwords)
}

func serialize(categories map[string]*Category, passwords map[int]*Password) [][]*Password {
	ret := [][]*Password{}
	cmap := map[string][]*Password{}
	for cname := range categories {
		ps := []*Password{}
		cmap[cname] = ps
	}
	for _, p := range passwords {
		ps := cmap[p.Category.ID]
		ps = append(ps, p)
		cmap[p.Category.ID] = ps
	}
	ifNil := func(sp *string) string {
		if sp == nil {
			return ""
		}
		return *sp
	}
	getCmpKey := func(p *Password) string {
		return fmt.Sprintf("%s-%s-%s-%s", p.Name, ifNil(p.User), ifNil(p.Password), ifNil(p.Mail))
	}
	for _, ps := range cmap {
		sort.Slice(ps, func(a int, b int) bool {
			p1 := ps[a]
			p2 := ps[b]
			return getCmpKey(p1) < getCmpKey(p2)
		})
		ret = append(ret, ps)
	}
	sort.Slice(ret, func(a int, b int) bool {
		p1 := ret[a]
		p2 := ret[b]
		return p1[0].Category.Order < p2[0].Category.Order
	})
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

func (p Password) String() string {
	return fmt.Sprintf("aaa: %v", p.Category)
	//return fmt.Sprintf(`{ id: %v, name: %v, desc: %v, category: %v, user: %v, password: %v, mail: %v, note: %v, created_at: %v, updated_at: %v }`,
	//	p.ID, p.Name, ifNil(p.Desc), p.Category, ifNil(p.User), ifNil(p.Password), ifNil(p.Mail), ifNil(p.Note), p.CreatedAt, p.UpdatedAt)
}
