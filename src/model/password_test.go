package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPasswordCreateRequest_Validate(t *testing.T) {
	category := Category{
		ID:    "category",
		Name:  "category name",
		Order: 10,
	}
	type fields struct {
		Name       string
		Desc       string
		CategoryID string
		User       string
		Password   string
		Mail       string
		Note       string
	}
	type args struct {
		cmap map[string]*Category
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		assertFunc func(password Password)
	}{
		{
			name: "success",
			fields: fields{
				Name:       "name",
				Desc:       "desc",
				CategoryID: "category",
				User:       "user",
				Password:   "password",
				Mail:       "mail",
				Note:       "note",
			},
			args: args{
				cmap: map[string]*Category{
					"category": &category,
				},
			},
			assertFunc: func(password Password) {
				assert.Equal(t, Password{
					Name:     "name",
					Desc:     "desc",
					Category: &category,
					User:     "user",
					Password: "password",
					Mail:     "mail",
					Note:     "note",
				}, password)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &PasswordCreateRequest{
				Name:       tt.fields.Name,
				Desc:       tt.fields.Desc,
				CategoryID: tt.fields.CategoryID,
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				Mail:       tt.fields.Mail,
				Note:       tt.fields.Note,
			}
			tt.assertFunc(*req.Validate(tt.args.cmap))
		})
	}
}

func strp(value string) *string {
	return &value
}

func TestPasswordUpdateRequest_Validate(t *testing.T) {
	category1 := Category{
		ID:    "category1",
		Name:  "name1",
		Order: 10,
	}
	category2 := Category{
		ID:    "category2",
		Name:  "name2",
		Order: 20,
	}
	type fields struct {
		Name       *string
		Desc       *string
		CategoryID *string
		User       *string
		Password   *string
		Mail       *string
		Note       *string
	}
	type args struct {
		pwd  *Password
		cmap map[string]*Category
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    assert.ErrorAssertionFunc
		assertFunc func(password Password)
	}{
		{
			name: "success",
			fields: fields{
				Name:       strp("name2"),
				Desc:       strp("desc2"),
				CategoryID: strp("category2"),
				User:       strp("user2"),
				Password:   strp("password2"),
				Mail:       strp("mail2"),
				Note:       strp("note2"),
			},
			args: args{
				pwd: &Password{
					Name:     "name1",
					Desc:     "desc1",
					Category: &category1,
					User:     "user1",
					Password: "password1",
					Mail:     "mail1",
					Note:     "note1",
				},
				cmap: map[string]*Category{
					"category1": &category1,
					"category2": &category2,
				},
			},
			wantErr: assert.NoError,
			assertFunc: func(password Password) {
				assert.Equal(t, Password{
					Name:     "name2",
					Desc:     "desc2",
					Category: &category2,
					User:     "user2",
					Password: "password2",
					Mail:     "mail2",
					Note:     "note2",
				}, password)
			},
		},
		{
			name: "invalid category",
			fields: fields{
				Name:       strp("name2"),
				Desc:       strp("desc2"),
				CategoryID: strp("category3"),
				User:       strp("user2"),
				Password:   strp("password2"),
				Mail:       strp("mail2"),
				Note:       strp("note2"),
			},
			args: args{
				pwd: &Password{
					Name:     "name1",
					Desc:     "desc1",
					Category: &category1,
					User:     "user1",
					Password: "password1",
					Mail:     "mail1",
					Note:     "note1",
				},
				cmap: map[string]*Category{
					"category1": &category1,
					"category2": &category2,
				},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &PasswordUpdateRequest{
				Name:       tt.fields.Name,
				Desc:       tt.fields.Desc,
				CategoryID: tt.fields.CategoryID,
				User:       tt.fields.User,
				Password:   tt.fields.Password,
				Mail:       tt.fields.Mail,
				Note:       tt.fields.Note,
			}
			tt.wantErr(t, req.Validate(tt.args.pwd, tt.args.cmap), fmt.Sprintf("Validate(%v, %v)", tt.args.pwd, tt.args.cmap))
			if tt.assertFunc != nil {
				tt.assertFunc(*tt.args.pwd)
			}
		})
	}
}
