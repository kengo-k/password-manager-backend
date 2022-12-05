package password

import (
	"bytes"
	"encoding/json"
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

type ApiCallWrapper func(method string, url string, req any, res any)

// get api function for test
func getGinHandler(api types.ApiCall) ApiCallWrapper {
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

	return func(method string, url string, req any, res any) {
		if req != nil {
			requestJson, _ := json.Marshal(req)
			request, _ := http.NewRequest(method, url, bytes.NewReader(requestJson))
			request.Header.Add("Content-Type", binding.MIMEJSON)
			ctx.Request = request
		}
		callApi(ctx)
		json.Unmarshal(response.Body.Bytes(), res)
	}
}

func TestGetPasswords(t *testing.T) {
	callApi := getGinHandler(GetPasswords)
	passwords := []model.Password{}
	callApi("GET", "/api/passwords", nil, &passwords)

	expectedLen := 9
	if len(passwords) != expectedLen {
		t.Errorf("got: %v, expected: %v", len(passwords), expectedLen)
	}
}
