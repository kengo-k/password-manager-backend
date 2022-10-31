package repo

import (
	"fmt"
	"strings"

	"github.com/kengo-k/password-manager/model"
)

type Repository interface {
	FindPasswords() []model.Password
	FindCategories() []model.Category
	SavePassword(p *model.Password)
	SaveCategory(cat *model.Category)
	DeletePassword(p *model.Password)
	DeleteCategory(cat *model.Category)
}

type RepositoryImpl struct {
	Passwords  map[int]*model.Password
	Categories map[string]*model.Category
}

func (r *RepositoryImpl) FindPasswords() []model.Password {
	return []model.Password{}
}

func (r *RepositoryImpl) FindCategories() []model.Category {
	var cats []model.Category
	for _, v := range r.Categories {
		cats = append(cats, *v)
	}
	return cats
}

func (r *RepositoryImpl) SavePassword(p *model.Password) {
	r.Passwords[p.ID] = p
}

func (r *RepositoryImpl) DeletePassword(p *model.Password) {
	delete(r.Passwords, p.ID)
}

func (r *RepositoryImpl) SaveCategory(cat *model.Category) {
	r.Categories[cat.Name] = cat
}

func (r *RepositoryImpl) DeleteCategory(cat *model.Category) {
	delete(r.Categories, cat.Name)
}

// 旧形式パスワードを読み込む
// ※移行が完了したら不要になる
func (r *RepositoryImpl) Init(mdLines []string) error {

	foundCategory := false
	foundHeader := false
	foundSeparator := false

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
	pid := 1
	for _, l := range mdLines {
		// 空行の場合はスキップする
		if len(l) == 0 {
			continue
		}
		var c model.Category
		// #で始まるコメント行の場合はカテゴリ名が記載されている
		if strings.HasPrefix(l, "#") {
			foundCategory = true
			_, categoryName, ok := strings.Cut(l, "#")
			if !ok {
				panic("failed to get category name")
			}
			c = model.Category{
				Name: categoryName,
				Desc: nil,
			}
			r.Categories[categoryName] = &c
			fmt.Printf("category: %s\n", categoryName)
			continue
		}
		if foundCategory {
			foundCategory = false
			foundHeader = true
			fmt.Printf("header: %s\n", l)
			continue
		}
		if foundHeader {
			foundHeader = false
			foundSeparator = true
			fmt.Println("separator")
			continue
		}
		if foundSeparator {
			columns := splitColumns(l)
			if len(columns) != 5 && len(columns) != 6 {
				panic(fmt.Sprintf("invalid column length: %v", len(columns)))
			}

			if len(columns) == 5 {
				p := &model.Password{
					ID:       pid,
					Name:     columns[0],
					Desc:     &columns[1],
					Category: c,
					User:     &columns[2],
					Password: &columns[3],
					Mail:     nil,
					Note:     &columns[4],
				}
				r.Passwords[p.ID] = p
				fmt.Printf("line: %v\n", *p)
			}
			if len(columns) == 6 {
				p := &model.Password{
					ID:       pid,
					Name:     columns[0],
					Desc:     &columns[1],
					Category: c,
					User:     &columns[2],
					Password: &columns[3],
					Mail:     &columns[4],
					Note:     &columns[5],
				}
				r.Passwords[p.ID] = p
				fmt.Printf("line: %v\n", *p)
			}
		}
	}

	return nil
}

func newRepositoryImpl() *RepositoryImpl {
	r := &RepositoryImpl{
		Categories: map[string]*model.Category{},
		Passwords:  map[int]*model.Password{},
	}
	return r
}
