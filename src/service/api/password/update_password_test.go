package password

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kengo-k/password-manager/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func strp(value string) *string {
	return &value
}

func TestUpdate(t *testing.T) {

	prepare := func(params gin.Params) IApiCallWrapper {
		return createApiWrapper(UpdatePassword, "PUT", "http://localhost/api/passwords/1", params)
	}

	type args struct {
		request model.PasswordUpdateRequest
	}
	tests := []struct {
		name        string
		args        args
		prepareFunc func() IApiCallWrapper
		assertFunc  func(code int, password model.Password)
	}{
		{
			name: "success",
			args: args{
				request: model.PasswordUpdateRequest{
					Name:       strp("new name1"),
					Desc:       strp("new desc1"),
					CategoryID: strp("cat2"),
					User:       strp("new user1"),
					Password:   strp("new password1"),
					Mail:       strp("new mail1"),
					Note:       strp("new note1"),
				},
			},
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "1"}})
			},
			assertFunc: func(code int, password model.Password) {
				assert.Equal(t, 200, code)
				assert.Equal(t, "new name1", password.Name)
			},
		},
		{
			name: "invalid id",
			args: args{
				request: model.PasswordUpdateRequest{},
			},
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "a"}})
			},
			assertFunc: func(code int, password model.Password) {
				assert.Equal(t, 400, code)
			},
		},
		{
			name: "failed to bind",
			args: args{
				request: model.PasswordUpdateRequest{},
			},
			prepareFunc: func() IApiCallWrapper {
				update := prepare(gin.Params{{Key: "id", Value: "a"}})
				update.SetMimeType(binding.MIMEXML)
				return update
			},
			assertFunc: func(code int, password model.Password) {
				assert.Equal(t, 400, code)
			},
		},
		{
			name: "password is not exists",
			args: args{
				request: model.PasswordUpdateRequest{},
			},
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "100"}})
			},
			assertFunc: func(code int, password model.Password) {
				assert.Equal(t, 400, code)
			},
		},
		{
			name: "category is not exists",
			args: args{
				request: model.PasswordUpdateRequest{
					Name:       strp("new name1"),
					Desc:       strp("new desc1"),
					CategoryID: strp("catXXX"),
					User:       strp("new user1"),
					Password:   strp("new password1"),
					Mail:       strp("new mail1"),
					Note:       strp("new note1"),
				},
			},
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "1"}})
			},
			assertFunc: func(code int, password model.Password) {
				assert.Equal(t, 400, code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			update := tt.prepareFunc()
			password := model.Password{}
			response, _ := update.CallApi(tt.args.request, &password)
			tt.assertFunc(response.Code, password)
		})
	}
}
