package model

import (
	"fmt"
	"time"
)

// password table model
type Password struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      *string   `json:"desc"`
	Category  *Category `json:"category"`
	User      *string   `json:"user"`
	Password  *string   `json:"password"`
	Mail      *string   `json:"mail"`
	Note      *string   `json:"note"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// request model for password creation
type PasswordCreateRequest struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	CategoryID string `json:"category_id"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Mail       string `json:"mail"`
	Note       string `json:"note"`
}

func (req *PasswordCreateRequest) Validate(cmap map[string]*Category) (*Password, error) {

	if len(req.Name) == 0 {
		return nil, fmt.Errorf("name is empty")
	}

	if len(req.CategoryID) == 0 {
		return nil, fmt.Errorf("category is empty")
	}

	cat, ok := cmap[req.CategoryID]
	if !ok {
		return nil, fmt.Errorf("category %v is not exists", req.CategoryID)
	}
	pwd := Password{
		Name:     req.Name,
		Desc:     &req.Desc,
		Category: cat,
		User:     &req.User,
		Password: &req.Password,
		Mail:     &req.Mail,
		Note:     &req.Note,
	}
	return &pwd, nil
}

// request model for password update
type PasswordUpdateRequest struct {
	Name       *string `json:"name"`
	Desc       *string `json:"desc"`
	CategoryID *string `json:"category_id"`
	User       *string `json:"user"`
	Password   *string `json:"password"`
	Mail       *string `json:"mail"`
	Note       *string `json:"note"`
}

func (req *PasswordUpdateRequest) Validate(pwd *Password, cmap map[string]*Category) error {
	setNewValue(req.Name, &pwd.Name)
	setNewValue(req.Desc, pwd.Desc)
	setNewValue(req.User, pwd.User)
	setNewValue(req.Password, pwd.Password)
	setNewValue(req.Mail, pwd.Mail)
	setNewValue(req.Note, pwd.Note)
	if req.CategoryID != nil {
		cat, ok := cmap[*req.CategoryID]
		if !ok {
			return fmt.Errorf("category %v is not exists", req.CategoryID)
		}
		pwd.Category = cat
	}
	return nil
}

func setNewValue(src *string, dest *string) {
	if src != nil {
		*dest = *src
	}
}
