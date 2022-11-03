package repo

import (
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
	for _, v := range r.database.Passwords {
		ret = append(ret, v)
	}
	return ret
}

func (r *Repository) FindCategories() []*model.Category {
	ret := make([]*model.Category, len(r.database.Categories))
	for _, v := range r.database.Categories {
		ret = append(ret, v)
	}
	return ret
}

func (r *Repository) SavePassword(p *model.Password) {
	r.database.Passwords[p.ID] = p
}

func (r *Repository) DeletePassword(p *model.Password) {
	delete(r.database.Passwords, p.ID)
}

func (r *Repository) SaveCategory(cat *model.Category) {
	r.database.Categories[cat.Name] = cat
}

func (r *Repository) DeleteCategory(cat *model.Category) {
	delete(r.database.Categories, cat.Name)
}
