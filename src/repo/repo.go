package repo

import (
	"fmt"

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

func ConvertMarkdown(passwords map[string][]*model.Password) []string {
	markdown := []string{}
	for categoryName, categoryPasswords := range passwords {
		head := categoryPasswords[0]
		categoryLine := fmt.Sprintf("# %s: %s", categoryName, *head.Category.Desc)
		headerLine := "|id|name|desc|user|password|mail|note|"
		separatorLine := "|---|---|---|---|---|---|---|"
		markdown = append(markdown, categoryLine)
		markdown = append(markdown, headerLine)
		markdown = append(markdown, separatorLine)
		for _, pwd := range categoryPasswords {
			line := fmt.Sprintf("|%v|%v|%v|%v|%v|%v|%v|",
				pwd.ID, pwd.Name, *pwd.Desc, *pwd.User, *pwd.Password, *pwd.Mail, *pwd.Note)
			markdown = append(markdown, line)
		}
	}
	return markdown
}
