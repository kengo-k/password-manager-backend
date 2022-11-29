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

func (r *Repository) GetPassword(id int) *model.Password {
	return r.database.PasswordMap[id]
}

func (r *Repository) GetNextPasswordId() int {
	return r.database.GetNextPasswordId()
}

func (r *Repository) GetCategories() map[string]*model.Category {
	return r.database.CategoryMap
}

func (r *Repository) SavePassword(p *model.Password) {
	r.database.PasswordMap[p.ID] = p
	pwds := r.database.CategorizedPasswords[p.Category.ID]
	pwds = append(pwds, p)
	r.database.CategorizedPasswords[p.Category.ID] = pwds
}

func (r *Repository) DeletePassword(p *model.Password) {
	delete(r.database.PasswordMap, p.ID)
}

func (r *Repository) SaveCategory(cat *model.Category) {
	r.database.CategoryMap[cat.ID] = cat
}

func (r *Repository) GetCategory(id string) *model.Category {
	return r.database.CategoryMap[id]
}

func (r *Repository) DeleteCategory(cat *model.Category) {
	delete(r.database.CategoryMap, cat.Name)
}
