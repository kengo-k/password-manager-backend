package password

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {

	prepare := func(params gin.Params) IApiCallWrapper {
		return createApiWrapper(DeletePassword, "DELETE", "http://localhost/api/passwords/1", params)
	}

	tests := []struct {
		name        string
		prepareFunc func() IApiCallWrapper
		assertFunc  func(code int)
	}{
		{
			name: "success",
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "1"}})
			},
			assertFunc: func(code int) {
				assert.Equal(t, 200, code)
			},
		},
		{
			name: "invalid id",
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "a"}})
			},
			assertFunc: func(code int) {
				assert.Equal(t, 400, code)
			},
		},
		{
			name: "password is not exists",
			prepareFunc: func() IApiCallWrapper {
				return prepare(gin.Params{{Key: "id", Value: "100"}})
			},
			assertFunc: func(code int) {
				assert.Equal(t, 404, code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delete := tt.prepareFunc()
			response, _ := delete.CallApi(nil, nil)
			tt.assertFunc(response.Code)
		})
	}
}
