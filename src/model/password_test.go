package model

import (
	"testing"
)

func getStrPointer(s string) *string {
	return &s
}

func TestApplyUpdateValues(t *testing.T) {
	pwd := Password{
		Name: "name",
		Desc: getStrPointer("desc"),
		Category: &Category{
			ID:   "cat1",
			Name: "category1",
		},
		User:     getStrPointer("user"),
		Password: getStrPointer("password"),
		Mail:     getStrPointer("mail"),
		Note:     getStrPointer("note"),
	}
	req := PasswordUpdateRequest{
		Name:       getStrPointer("new name"),
		Desc:       getStrPointer("new desc"),
		CategoryID: getStrPointer("cat2"),
		User:       getStrPointer("new user"),
		Password:   getStrPointer("new password"),
		Mail:       getStrPointer("new mail"),
		Note:       getStrPointer("new note"),
	}
	req.Validate(&pwd, map[string]*Category{
		"cat1": {
			ID:   "cat1",
			Name: "category1",
		},
		"cat2": {
			ID:   "cat2",
			Name: "category2",
		},
	})
	if pwd.Name != "new name" {
		t.Errorf("got: %v, expected: %v", pwd.Name, "new name")
	}
	if *pwd.Desc != "new desc" {
		t.Errorf("got: %v, expected: %v", *pwd.Desc, "new desc")
	}
	if pwd.Category.ID != "cat2" {
		t.Errorf("got: %v, expected: %v", pwd.Category.ID, "cat2")
	}
	if *pwd.User != "new user" {
		t.Errorf("got: %v, expected: %v", *pwd.User, "new user")
	}
	if *pwd.Password != "new password" {
		t.Errorf("got: %v, expected: %v", *pwd.Password, "new password")
	}
	if *pwd.Mail != "new mail" {
		t.Errorf("got: %v, expected: %v", *pwd.Mail, "new mail")
	}
	if *pwd.Note != "new note" {
		t.Errorf("got: %v, expected: %v", *pwd.Note, "new note")
	}
}
