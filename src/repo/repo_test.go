package repo

import (
	"github.com/kengo-k/password-manager/model"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		database *model.Database
	}
	tests := []struct {
		name string
		args args
		want func(args args, repo *Repository)
	}{
		{
			name: "success",
			args: args{},
			want: func(args args, repo *Repository) {
				assert.Equal(t, repo.database, args.database)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepository(tt.args.database)
			tt.want(tt.args, got)
		})
	}
}

func TestRepository_DeleteCategory(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		cat *model.Category
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			repo.DeleteCategory(tt.args.cat)
		})
	}
}

func TestRepository_DeletePassword(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		p *model.Password
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			repo.DeletePassword(tt.args.p)
		})
	}
}

func TestRepository_FindPasswords(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	tests := []struct {
		name   string
		fields fields
		want   []*model.Password
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.FindPasswords(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetCategories(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]*model.Category
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.GetCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetCategory(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Category
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.GetCategory(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetNextPasswordId(t *testing.T) {
	tests := []struct {
		name           string
		createDatabase func() *model.Database
		want           int
	}{
		{
			name: "success",
			createDatabase: func() *model.Database {
				return model.NewDatabase()
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := tt.createDatabase()
			repo := &Repository{
				database: database,
			}
			if got := repo.GetNextPasswordId(); got != tt.want {
				t.Errorf("GetNextPasswordId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetPassword(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Password
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.GetPassword(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetSortedCategories(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	tests := []struct {
		name   string
		fields fields
		want   []*model.Category
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.GetSortedCategories(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSortedCategories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsDirty(t *testing.T) {
	tests := []struct {
		name           string
		createDatabase func() *model.Database
		want           bool
	}{
		{
			name: "is clean",
			createDatabase: func() *model.Database {
				return model.NewDatabase()
			},
			want: false,
		},
		{
			name: "is dirty",
			createDatabase: func() *model.Database {
				database := model.NewDatabase()
				database.SetDirty(true)
				return database
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := tt.createDatabase()
			repo := &Repository{
				database: database,
			}
			if got := repo.IsDirty(); got != tt.want {
				t.Errorf("IsDirty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SaveCategory(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		cat *model.Category
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			repo.SaveCategory(tt.args.cat)
		})
	}
}

func TestRepository_SavePassword(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	type args struct {
		p *model.Password
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			repo.SavePassword(tt.args.p)
		})
	}
}

func TestRepository_Serialize(t *testing.T) {
	type fields struct {
		database *model.Database
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]*model.Password
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				database: tt.fields.database,
			}
			if got := repo.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetClean(t *testing.T) {
	tests := []struct {
		name           string
		createDatabase func() *model.Database
		beforeAssert   func(repo *Repository)
		afterAssert    func(repo *Repository)
	}{
		{
			name: "success",
			createDatabase: func() *model.Database {
				database := model.NewDatabase()
				database.SetDirty(true)
				return database
			},
			beforeAssert: func(repo *Repository) {
				assert.Equal(t, true, repo.IsDirty())
			},
			afterAssert: func(repo *Repository) {
				assert.Equal(t, false, repo.IsDirty())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database := tt.createDatabase()
			repo := &Repository{
				database: database,
			}
			tt.beforeAssert(repo)
			repo.SetClean()
			tt.afterAssert(repo)
		})
	}
}
