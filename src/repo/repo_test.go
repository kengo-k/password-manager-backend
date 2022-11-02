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
				Name:     "name1",
				Desc:     p("desc1"),
				Category: cate1,
				User:     p("user1"),
				Password: p("password1"),
				Mail:     p("mail1"),
				Note:     p("note1"),
			},
			{
				ID:       2,
				Name:     "name2",
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
				Name:     "name3",
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
	gotLen := len(markdown)
	expectedLen := 9
	if gotLen != expectedLen {
		t.Errorf("length: got=%v, expected:%v", gotLen, expectedLen)
	}

	exptects := []string{
		"# cate1: desc1",
		"|id|name|desc|user|password|mail|note|",
		"|---|---|---|---|---|---|---|",
		"|1|name1|desc1|user1|password1|mail1|note1|",
		"|2|name2|desc2|user2|password2|mail2|note2|",
		"# cate2: desc2",
		"|id|name|desc|user|password|mail|note|",
		"|---|---|---|---|---|---|---|",
		"|3|name3|desc3|user3|password3|mail3|note3|",
	}

	for i, expect := range exptects {
		got := markdown[i]
		if got != expect {
			t.Errorf("line[%v]: got=%v, expected=%v", i, got, expect)
		}
	}

}
