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

func (repo *Repository) GetNextPasswordId() int {
	return repo.database.GetNextPasswordId()
}

func (repo *Repository) IsDirty() bool {
	return repo.database.IsDirty()
}

func (repo *Repository) SetClean() {
	repo.database.SetDirty(false)
}

// get single password by id
func (repo *Repository) GetPassword(id int) *model.Password {
	return repo.database.PasswordMap[id]
}

// find passwords by search condition
func (repo *Repository) FindPasswords() []*model.Password {
	ret := []*model.Password{}
	sortedCats := repo.GetSortedCategories()
	for _, cat := range sortedCats {
		pwds := repo.database.CategorizedPasswords[cat.ID]
		sort.SliceStable(pwds, func(i, j int) bool {
			a := pwds[i]
			b := pwds[j]
			return a.Name < b.Name
		})
		ret = append(ret, pwds...)
	}
	return ret
}

// save password into database
func (repo *Repository) SavePassword(p *model.Password) {
	if _, ok := repo.database.PasswordMap[p.ID]; !ok {
		pwds := repo.database.CategorizedPasswords[p.Category.ID]
		pwds = append(pwds, p)
		repo.database.CategorizedPasswords[p.Category.ID] = pwds
	}
	repo.database.PasswordMap[p.ID] = p
	repo.database.SetDirty(true)
}

// delete password from database
func (repo *Repository) DeletePassword(p *model.Password) {
	delete(repo.database.PasswordMap, p.ID)
	repo.database.SetDirty(true)
}

// get single category by id
func (repo *Repository) GetCategory(id string) *model.Category {
	return repo.database.CategoryMap[id]
}

// get all categories
func (repo *Repository) GetCategories() map[string]*model.Category {
	return repo.database.CategoryMap
}

// get sorted all categories
func (repo *Repository) GetSortedCategories() []*model.Category {
	ret := []*model.Category{}
	for _, cat := range repo.database.CategoryMap {
		ret = append(ret, cat)
	}
	sort.SliceStable(ret, func(i, j int) bool {
		a := ret[i]
		b := ret[j]
		return a.Order < b.Order
	})
	return ret
}

// save category into database
func (repo *Repository) SaveCategory(cat *model.Category) {
	repo.database.CategoryMap[cat.ID] = cat
	repo.database.SetDirty(true)
}

// delete category from database
func (repo *Repository) DeleteCategory(cat *model.Category) {
	// TODO カテゴリ配下のパスワードを一括で消すorカテゴリ配下にパスワードが存在する場合はエラーにする
	delete(repo.database.CategoryMap, cat.Name)
	repo.database.SetDirty(true)
}

func (repo *Repository) Serialize() [][]*model.Password {
	return repo.database.Serialize()
}
