package repo

import (
	"sort"

	"github.com/kengo-k/password-manager/model"
)

type Repository struct {
	database *model.Database
}

func NewRepository(database *model.Database) *Repository {
	r := &Repository{
		database: database,
	}
	return r
}

func (r *Repository) FindPasswords() []*model.Password {
	ret := []*model.Password{}
	sortedCats := r.database.GetSortedCategories()
	for _, cat := range sortedCats {
		pwds := r.database.CategorizedPasswords[cat.ID]
		sort.SliceStable(pwds, func(i, j int) bool {
			a := pwds[i]
			b := pwds[j]
			return a.Name < b.Name
		})
		ret = append(ret, pwds...)
	}
	return ret
}

func (r *Repository) FindCategories() []*model.Category {
	ret := []*model.Category{}
	for _, v := range r.database.CategoryMap {
		ret = append(ret, v)
	}
	return ret
}

func (r *Repository) SavePassword(p *model.Password) {
	r.database.PasswordMap[p.ID] = p
}

func (r *Repository) DeletePassword(p *model.Password) {
	delete(r.database.PasswordMap, p.ID)
}

func (r *Repository) SaveCategory(cat *model.Category) {
	r.database.CategoryMap[cat.Name] = cat
}

func (r *Repository) DeleteCategory(cat *model.Category) {
	delete(r.database.CategoryMap, cat.Name)
}
