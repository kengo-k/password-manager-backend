package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNextPasswordId(t *testing.T) {
	database := NewDatabase()
	assert.Equal(t, 1, database.GetNextPasswordId())
	assert.Equal(t, 2, database.GetNextPasswordId())
}

func TestDirty(t *testing.T) {
	database := NewDatabase()
	assert.Equal(t, false, database.IsDirty())
	database.SetDirty(true)
	assert.Equal(t, true, database.IsDirty())
}

func TestNewDatabase(t *testing.T) {
	database := NewDatabase()
	assert.NotNil(t, database)
}

func TestSplitColumns(t *testing.T) {
	lines := splitColumns("|aaa|bbb|ccc|")
	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "aaa", lines[0])
	assert.Equal(t, "bbb", lines[1])
	assert.Equal(t, "ccc", lines[2])
}

func TestGetCategory(t *testing.T) {
	type testSetting struct {
		line       string
		want       assert.ValueAssertionFunc
		assertFunc func(c Category)
		wantError  assert.ErrorAssertionFunc
	}

	tss := []testSetting{
		{
			line:      "---",
			want:      assert.Nil,
			wantError: assert.Error,
		},
		{
			line:      "# hello",
			want:      assert.Nil,
			wantError: assert.Error,
		},
		{
			line:      "# hello: ",
			want:      assert.Nil,
			wantError: assert.Error,
		},
		{
			line:      "# hello: name,order",
			want:      assert.Nil,
			wantError: assert.Error,
		},
		{
			line:      "# hello: name=hello,order=aaa",
			want:      assert.Nil,
			wantError: assert.Error,
		},
		{
			line: "# hello: name=hello,order=1",
			assertFunc: func(c Category) {
				assert.Equal(t, c, Category{Name: "hello", ID: "hello", Order: 1})
			},
			wantError: assert.NoError,
		},
	}

	for _, ts := range tss {
		v, err := getCategory(ts.line)
		if ts.want != nil {
			assert.True(t, ts.want(t, v))
		}
		if ts.assertFunc != nil {
			ts.assertFunc(*v)
		}
		assert.True(t, ts.wantError(t, err))
	}
}

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
