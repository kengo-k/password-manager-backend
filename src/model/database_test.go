package model

import (
	"fmt"
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

func Test_lineContext_SetCategoryOn(t *testing.T) {
	type fields struct {
		isCategory bool
		isHeader   bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				isCategory: false,
				isHeader:   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &lineContext{
				isCategory: tt.fields.isCategory,
				isHeader:   tt.fields.isHeader,
			}
			lc.SetCategoryOn()
			assert.Equal(t, true, lc.isCategory)
		})
	}
}

func Test_lineContext_ShouldSkip(t *testing.T) {
	type fields struct {
		isCategory bool
		isHeader   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "false, false",
			fields: fields{
				isCategory: false,
				isHeader:   false,
			},
			want: false,
		},
		{
			name: "true, false",
			fields: fields{
				isCategory: true,
				isHeader:   false,
			},
			want: true,
		},
		{
			name: "false, true",
			fields: fields{
				isCategory: false,
				isHeader:   true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &lineContext{
				isCategory: tt.fields.isCategory,
				isHeader:   tt.fields.isHeader,
			}
			assert.Equalf(t, tt.want, lc.ShouldSkip(), "ShouldSkip()")
		})
	}
}

func TestDatabase_Init(t *testing.T) {
	type fields struct {
		dirty                bool
		maxPasswordId        int
		PasswordMap          map[int]*Password
		CategoryMap          map[string]*Category
		CategorizedPasswords map[string][]*Password
	}
	type args struct {
		mdLines []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:   "invalid category line",
			fields: fields{},
			args: args{
				mdLines: []string{
					"# invalid category line",
				},
			},
			wantErr: assert.Error,
		}, {
			name: "invalid column size",
			fields: fields{
				PasswordMap:          map[int]*Password{},
				CategoryMap:          map[string]*Category{},
				CategorizedPasswords: map[string][]*Password{},
			},
			args: args{
				mdLines: []string{
					"# cat1: name=category1,order=10",
					"| 名称 | 説明 | ユーザ | パスワード | メール | 備考 |",
					"|---|---|---|---|---|---|---|",
					"|col1|col2|col3|col4|col5|col6|col7|",
				},
			},
			wantErr: assert.Error,
		}, {
			name: "success",
			fields: fields{
				PasswordMap:          map[int]*Password{},
				CategoryMap:          map[string]*Category{},
				CategorizedPasswords: map[string][]*Password{},
			},
			args: args{
				mdLines: []string{
					"# cat1: name=category1,order=1",
					"| 名称 | 説明 | ユーザ | パスワード | メール | 備考 |",
					"|---|---|---|---|---|---|",
					"|col1|col2|col3|col4|col5|col6|",
					"",
					"# cat2: name=category2,order=2",
					"| 名称 | 説明 | ユーザ | パスワード | メール | 備考 |",
					"|---|---|---|---|---|---|",
					"|col-a|col-b|col-c|col-d|col-e|col-f|",
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				dirty:                tt.fields.dirty,
				maxPasswordId:        tt.fields.maxPasswordId,
				PasswordMap:          tt.fields.PasswordMap,
				CategoryMap:          tt.fields.CategoryMap,
				CategorizedPasswords: tt.fields.CategorizedPasswords,
			}
			tt.wantErr(t, d.Init(tt.args.mdLines), fmt.Sprintf("Init(%v)", tt.args.mdLines))
		})
	}
}

func TestDatabase_Serialize(t *testing.T) {
	type fields struct {
		dirty                bool
		maxPasswordId        int
		PasswordMap          map[int]*Password
		CategoryMap          map[string]*Category
		CategorizedPasswords map[string][]*Password
	}
	cat1 := Category{
		ID:    "cat1",
		Name:  "category1",
		Order: 10,
	}
	cat2 := Category{
		ID:    "cat2",
		Name:  "category2",
		Order: 20,
	}
	pwd1 := Password{
		ID:       1,
		Name:     "name1",
		Desc:     "desc1",
		Category: &cat1,
		User:     "user1",
		Password: "password1",
		Mail:     "mail1",
		Note:     "note1",
	}
	pwd2 := Password{
		ID:       2,
		Name:     "name2",
		Desc:     "desc2",
		Category: &cat2,
		User:     "user2",
		Password: "password2",
		Mail:     "mail2",
		Note:     "note2",
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]*Password
	}{
		{
			name: "success",
			fields: fields{
				PasswordMap: map[int]*Password{
					1: &pwd1,
					2: &pwd2,
				},
				CategoryMap: map[string]*Category{
					"cat1": &cat1,
					"cat2": &cat2,
				},
				CategorizedPasswords: map[string][]*Password{
					"cat1": {&pwd1},
					"cat2": {&pwd2},
				},
			},
			want: [][]*Password{{&pwd1}, {&pwd2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				dirty:                tt.fields.dirty,
				maxPasswordId:        tt.fields.maxPasswordId,
				PasswordMap:          tt.fields.PasswordMap,
				CategoryMap:          tt.fields.CategoryMap,
				CategorizedPasswords: tt.fields.CategorizedPasswords,
			}
			assert.Equalf(t, tt.want, d.Serialize(), "Serialize()")
		})
	}
}
