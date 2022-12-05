package model

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Database struct {
	dirty                bool
	maxPasswordId        int
	PasswordMap          map[int]*Password
	CategoryMap          map[string]*Category
	CategorizedPasswords map[string][]*Password
}

func (d *Database) GetNextPasswordId() int {
	d.maxPasswordId++
	return d.maxPasswordId
}

func (d *Database) SetDirty(dirty bool) {
	d.dirty = dirty
}

func (d *Database) IsDirty() bool {
	return d.dirty
}

func NewDatabase() *Database {
	return &Database{
		dirty:                false,
		PasswordMap:          map[int]*Password{},
		CategoryMap:          map[string]*Category{},
		CategorizedPasswords: map[string][]*Password{},
	}
}

func splitColumns(line string) []string {
	ret := []string{}
	columns := strings.Split(line, "|")
	for i, column := range columns {
		if i == 0 || i == len(columns)-1 {
			continue
		}
		column = strings.TrimSpace(column)
		ret = append(ret, column)
	}
	return ret
}

func getCategory(l string) (*Category, error) {
	_, categoryLine, ok := strings.Cut(l, "#")
	if !ok {
		return nil, fmt.Errorf("failed to get catgory line")
	}
	categoryLine = strings.TrimSpace(categoryLine)
	categoryId, attrLine, ok := strings.Cut(categoryLine, ":")
	if !ok {
		return nil, fmt.Errorf("failed to get categoryId")
	}
	attrLine = strings.TrimSpace(attrLine)
	attrs := strings.Split(attrLine, ",")
	attrMap := map[string]string{}
	for _, attr := range attrs {
		k, v, ok := strings.Cut(attr, "=")
		if !ok {
			return nil, fmt.Errorf("failed to parse category attribute")
		}
		attrMap[k] = v
	}
	orderVal, err := strconv.ParseInt(attrMap["order"], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse category order to number")
	}
	category := Category{
		ID:    categoryId,
		Name:  attrMap["name"],
		Order: int(orderVal),
	}
	return &category, nil
}

type lineContext struct {
	isCategory bool
	isHeader   bool
}

func (lc *lineContext) SetCategoryOn() {
	lc.isCategory = true
}

func (lc *lineContext) ShouldSkip() bool {
	if lc.isCategory {
		lc.isCategory = false
		lc.isHeader = true
		return true
	}
	if lc.isHeader {
		lc.isHeader = false
		return true
	}
	return false
}

func (d *Database) Init(mdLines []string) error {

	lineCtx := lineContext{}

	var cat *Category
	var err error

	for _, l := range mdLines {
		// skip empty line
		if len(l) == 0 {
			continue
		}

		// line starts with `#` has category info
		if strings.HasPrefix(l, "#") {
			lineCtx.SetCategoryOn()
			cat, err = getCategory(l)
			if err != nil {
				return fmt.Errorf("failed to get category")
			}
			d.CategoryMap[cat.ID] = cat
			d.CategorizedPasswords[cat.ID] = []*Password{}
			continue
		}

		// if line is header or separator, skip the line
		if lineCtx.ShouldSkip() {
			continue
		}

		columns := splitColumns(l)
		if len(columns) != 6 {
			return fmt.Errorf("faild to load, invalid column length: %v", len(columns))
		}

		p := &Password{
			ID:       d.GetNextPasswordId(),
			Name:     columns[0],
			Desc:     columns[1],
			Category: cat,
			User:     columns[2],
			Password: columns[3],
			Mail:     columns[4],
			Note:     columns[5],
		}
		d.PasswordMap[p.ID] = p
		pwds := d.CategorizedPasswords[cat.ID]
		pwds = append(pwds, p)
		d.CategorizedPasswords[cat.ID] = pwds
	}

	return nil
}

func (d *Database) Serialize() [][]*Password {
	return serialize(d.CategoryMap, d.PasswordMap)
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
	getCmpKey := func(p *Password) string {
		return fmt.Sprintf("%s-%s-%s-%s-%s", p.Name, p.Desc, p.User, p.Password, p.Mail)
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
