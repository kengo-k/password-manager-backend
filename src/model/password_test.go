package model

import (
	"testing"
)

func getStrPointer(s string) *string {
	return &s
}

func TestApplyUpdateValues(t *testing.T) {
	pwd := Password{
		Name:     "name",
		Desc:     getStrPointer("desc"),
		User:     getStrPointer("user"),
		Password: getStrPointer("password"),
		Mail:     getStrPointer("mail"),
		Note:     getStrPointer("note"),
	}
	req := PasswordUpdateRequest{
		Name:     getStrPointer("new name"),
		Desc:     getStrPointer("new desc"),
		User:     getStrPointer("new user"),
		Password: getStrPointer("new password"),
		Mail:     getStrPointer("new mail"),
		Note:     getStrPointer("new note"),
	}
	req.ApplyValuesWithoutCategory(&pwd)
	if pwd.Name != "new name" {
		t.Errorf("got: %v, expected: %v", pwd.Name, "new name")
	}
	if *pwd.Desc != "new desc" {
		t.Errorf("got: %v, expected: %v", *pwd.Desc, "new desc")
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
