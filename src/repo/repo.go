package repo

import (
	"fmt"
	"strings"

	"github.com/kengo-k/password-manager/git"
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
	Passwords  map[string]*model.Password
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
	r.Passwords[*p.ID] = p
}

func (r *RepositoryImpl) DeletePassword(p *model.Password) {

}

func (r *RepositoryImpl) SaveCategory(cat *model.Category) {
	r.Categories[*cat.ID] = cat
}

func (r *RepositoryImpl) DeleteCategory(cat *model.Category) {
}

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

	for _, l := range mdLines {
		// 空行の場合はスキップする
		if len(l) == 0 {
			continue
		}
		// #で始まるコメント行の場合はカテゴリ名が記載されている
		if strings.HasPrefix(l, "#") {
			foundCategory = true
			_, category, ok := strings.Cut(l, "#")
			if !ok {
				panic("failed to get category name")
			}
			fmt.Printf("category: %s\n", category)
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
			if len(columns) != 7 {
				panic("invalid column length")
			}
			p := model.Password{
				ID:        &columns[0],
				User:      &columns[1],
				Password:  &columns[2],
				Mail:      &columns[3],
				Note:      &columns[4],
				CreatedAt: nil,
				UpdatedAt: nil,
			}
			fmt.Printf("line: %v\n", p)
		}
	}

	return nil
}

func NewRepository() Repository {
	g := &git.Git{}
	passwords, err := g.Checkout()
	if err != nil {
		panic("failed to checkout passwords")
	}
	r := &RepositoryImpl{}
	r.Init(passwords)
	return r
}
