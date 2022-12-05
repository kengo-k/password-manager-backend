package password

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

type ApiCallWrapper func(method string, url string, req any, res any) error

// get api function for test
func createApiWrapper(api types.ApiCall) ApiCallWrapper {
	gin.SetMode(gin.TestMode)
	// get context
	config := env.NewConfig("testdata/.test.env")
	context := context.NewContext(runmode.FILE_TO_FILE, config)
	// init database
	passwords, _ := context.Load()
	database := model.NewDatabase()
	database.Init(passwords)
	// init repository
	repo := repo.NewRepository(database)
	// get api function
	callApi := api(repo, context)
	// init response and context
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)

	return func(method string, url string, req any, res any) error {
		// if req exists, set request params to context
		if req != nil {
			requestJson, err := json.Marshal(req)
			if err != nil {
				return err
			}
			request, err := http.NewRequest(method, url, bytes.NewReader(requestJson))
			if err != nil {
				return err
			}
			request.Header.Add("Content-Type", binding.MIMEJSON)
			ctx.Request = request
		}
		callApi(ctx)
		return json.Unmarshal(response.Body.Bytes(), res)
	}
}

func TestGetPasswords(t *testing.T) {
	callApi := createApiWrapper(GetPasswords)

	passwords := []model.Password{}
	err := callApi("GET", "/api/passwords", nil, &passwords)
	if err != nil {
		t.Errorf(fmt.Sprintf("failed to call api: %v", err))
	}

	expectedLen := 9
	if len(passwords) != expectedLen {
		t.Errorf("got: %v, expected: %v", len(passwords), expectedLen)
	}
}
