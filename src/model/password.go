package model

import "time"

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
	Name       *string `json:"name"`
	Desc       *string `json:"desc"`
	CategoryID *string `json:"category_id"`
	User       *string `json:"user"`
	Password   *string `json:"password"`
	Mail       *string `json:"mail"`
	Note       *string `json:"note"`
}

// request model for password update
type PasswordUpdateRequest struct {
	ID         int     `json:"id"`
	Name       *string `json:"name"`
	Desc       *string `json:"desc"`
	CategoryID *string `json:"category_id"`
	User       *string `json:"user"`
	Password   *string `json:"password"`
	Mail       *string `json:"mail"`
	Note       *string `json:"note"`
}

// set new value to password without category
func (pwd *Password) ApplyUpdateValues(req *PasswordUpdateRequest) {
	setNewValue(req.Name, &pwd.Name)
	setNewValue(req.Desc, pwd.Desc)
	setNewValue(req.User, pwd.User)
	setNewValue(req.Password, pwd.Password)
	setNewValue(req.Mail, pwd.Mail)
	setNewValue(req.Note, pwd.Note)
}

func setNewValue(src *string, dest *string) {
	if src != nil {
		*dest = *src
	}
}
