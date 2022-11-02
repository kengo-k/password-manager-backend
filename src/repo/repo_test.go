package repo

import (
	"testing"

	"github.com/kengo-k/password-manager/git"
	"github.com/kengo-k/password-manager/model"
)

func TestInitRepository(t *testing.T) {
	g := &git.Git{}
	r := newRepositoryImpl()
	passwords, err := g.Checkout()
	if err != nil {
		t.Errorf("failed to checkout passwords: %v", err)
	}
	err = r.Init(passwords)
	if err != nil {
		t.Errorf("failed to init repository: %v", err)
	}
}

func TestConvertMarkdown(t *testing.T) {
	p := func(s string) *string {
		return &s
	}
	cate1 := model.Category{
		Name: "cate1",
		Desc: p("desc1"),
	}
	cate2 := model.Category{
		Name: "cate2",
		Desc: p("desc2"),
	}
	input := map[string][]*model.Password{
		"cate1": {
			{
				ID:       1,
				Name:     "item1",
				Desc:     p("desc1"),
				Category: cate1,
				User:     p("user1"),
				Password: p("password1"),
				Mail:     p("mail1"),
				Note:     p("note1"),
			},
			{
				ID:       2,
				Name:     "item2",
				Desc:     p("desc2"),
				Category: cate1,
				User:     p("user2"),
				Password: p("password2"),
				Mail:     p("mail2"),
				Note:     p("note2"),
			},
		},
		"cate2": {
			{
				ID:       3,
				Name:     "item3",
				Desc:     p("desc3"),
				Category: cate2,
				User:     p("user3"),
				Password: p("password3"),
				Mail:     p("mail3"),
				Note:     p("note3"),
			},
		},
	}
	markdown := ConvertMarkdown(input)
	got := len(markdown)
	expected := 9
	if got != expected {
		t.Errorf("length: got=%v, expected:%v", got, expected)
	}
}
