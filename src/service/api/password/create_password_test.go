package password

import (
	"github.com/gin-gonic/gin"
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/kengo-k/password-manager/model"
	"github.com/stretchr/testify/assert"
)

func TestCreatePassword(t *testing.T) {

	prepare := func() IApiCallWrapper {
		return createApiWrapper(CreatePassword, "POST", "/api/passwords", gin.Params{})
	}

	type testSetting struct {
		name        string
		request     model.PasswordCreateRequest
		prepareFunc func() IApiCallWrapper
		assertFunc  func(ts testSetting, code int, password model.Password)
	}

	tss := []testSetting{
		{
			name: "success",
			request: model.PasswordCreateRequest{
				Name:       "nameX",
				Desc:       "descX",
				CategoryID: "cat2",
			},
			prepareFunc: prepare,
			assertFunc: func(testItem testSetting, code int, password model.Password) {
				assert.Equal(t, code, 201)
				assert.Equal(t, testItem.request.Name, password.Name)
				assert.Equal(t, testItem.request.Desc, password.Desc)
			},
		},
		{
			name:    "failed to bind",
			request: model.PasswordCreateRequest{},
			prepareFunc: func() IApiCallWrapper {
				create := prepare()
				// invalid mimetype
				create.SetMimeType(binding.MIMEXML)
				return create
			},
			assertFunc: func(testItem testSetting, code int, password model.Password) {
				assert.Equal(t, code, 400)
			},
		},
		{
			name: "invalid category",
			request: model.PasswordCreateRequest{
				Name: "nameX",
				Desc: "descX",
				// invalid category id
				CategoryID: "cat999",
			},
			prepareFunc: prepare,
			assertFunc: func(testItem testSetting, code int, password model.Password) {
				assert.Equal(t, code, 400)
			},
		},
	}

	for _, ts := range tss {
		create := ts.prepareFunc()
		password := model.Password{}
		response, _ := create.CallApi(ts.request, &password)
		ts.assertFunc(ts, response.Code, password)
	}
}
