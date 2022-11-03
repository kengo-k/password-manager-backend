package model

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	p := func(s string) *string {
		return &s
	}
	cat := map[string]*Category{
		"Cat3": {Name: "Cat3", Desc: p("Desc3")},
		"Cat7": {Name: "Cat7", Desc: p("Desc7")},
		"Cat1": {Name: "Cat1", Desc: p("Desc1")},
		"Cat5": {Name: "Cat5", Desc: p("Desc5")},
		"Cat2": {Name: "Cat2", Desc: p("Desc2")},
	}
	pwd := map[int]*Password{
		100: {Name: "name100", Category: *cat["Cat3"]},
		101: {Name: "name101", Category: *cat["Cat3"]},
		102: {Name: "name102", Category: *cat["Cat3"]},

		200: {Name: "name200", Category: *cat["Cat7"]},
		201: {Name: "name201", Category: *cat["Cat7"]},

		300: {Name: "name300", Category: *cat["Cat1"]},

		400: {Name: "name400", Category: *cat["Cat5"]},
		401: {Name: "name401", Category: *cat["Cat5"]},

		500: {Name: "name500", Category: *cat["Cat2"]},
		501: {Name: "name501", Category: *cat["Cat2"]},
	}
	serialized := serialize(cat, pwd)
	categorySize := len(serialized)
	expCatSize := 5
	if categorySize != expCatSize {
		t.Errorf("category size: got=%v, expected=%v", categorySize, expCatSize)
	}
	type Table struct {
		Len int
		Cat string
	}
	table := []Table{
		{Len: 1, Cat: "Cat1"},
		{Len: 2, Cat: "Cat2"},
		{Len: 3, Cat: "Cat3"},
		{Len: 2, Cat: "Cat5"},
		{Len: 2, Cat: "Cat7"},
	}
	for i, passwords := range serialized {
		if len(passwords) == 0 {
			t.Errorf("len[%v] is 0", i)
		}
		tbl := table[i]
		p := passwords[0]
		gotLen := len(passwords)
		if gotLen != tbl.Len {
			t.Errorf("length: got=%v, expected:%v", gotLen, tbl.Len)
		}
		if p.Category.Name != tbl.Cat {
			t.Errorf("category: got=%v, expected:%v", p.Category.Name, tbl.Cat)
		}
	}
}
