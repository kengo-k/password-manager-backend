package model

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	cat := map[string]*Category{
		"mail":    {ID: "mail", Name: "MAIL", Order: 30},
		"tech":    {ID: "tech", Name: "TECH", Order: 51},
		"money":   {ID: "money", Name: "MONEY", Order: 2},
		"private": {ID: "private", Name: "PRIVATE", Order: 40},
		"other":   {ID: "other", Name: "OTHER", Order: 5},
	}
	pwd := map[int]*Password{
		100: {Name: "name100", Category: cat["mail"]},
		101: {Name: "name101", Category: cat["mail"]},
		102: {Name: "name102", Category: cat["mail"]},

		200: {Name: "name200", Category: cat["tech"]},
		201: {Name: "name201", Category: cat["tech"]},

		300: {Name: "name300", Category: cat["money"]},

		400: {Name: "name400", Category: cat["private"]},
		401: {Name: "name401", Category: cat["private"]},

		500: {Name: "name500", Category: cat["other"]},
		501: {Name: "name501", Category: cat["other"]},
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
		{Len: 1, Cat: "MONEY"},
		{Len: 2, Cat: "OTHER"},
		{Len: 3, Cat: "MAIL"},
		{Len: 2, Cat: "PRIVATE"},
		{Len: 2, Cat: "TECH"},
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
