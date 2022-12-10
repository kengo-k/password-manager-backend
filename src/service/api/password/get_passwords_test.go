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

type IApiCallWrapper interface {
	CallApi(req any, res any) (*httptest.ResponseRecorder, error)
	SetMimeType(mimeType string)
	SetRepository(repo *repo.Repository)
}

type ApiCallWrapper struct {
	mimeType string
	callApi  func(req any, res any) (*httptest.ResponseRecorder, error)
	repo     *repo.Repository
}

func (wrapper *ApiCallWrapper) CallApi(req any, res any) (*httptest.ResponseRecorder, error) {
	return wrapper.callApi(req, res)
}

func (wrapper *ApiCallWrapper) SetMimeType(mimeType string) {
	wrapper.mimeType = mimeType
}

func (wrapper *ApiCallWrapper) SetRepository(repo *repo.Repository) {
	wrapper.repo = repo
}

func newTestRepository() *repo.Repository {
	// get context
	config := env.NewConfig("testdata/.test.env")
	context := context.NewContext(runmode.FILE_TO_FILE, config)
	// init database
	passwords, _ := context.Load()
	database := model.NewDatabase()
	database.Init(passwords)
	// init repository
	repo := repo.NewRepository(database)
	return repo
}

// get api function for test
func createApiWrapper(api types.ApiCall, method string, url string, params gin.Params) IApiCallWrapper {
	gin.SetMode(gin.TestMode)
	// get context
	config := env.NewConfig("testdata/.test.env")
	context := context.NewContext(runmode.FILE_TO_FILE, config)
	// init repository
	repo := newTestRepository()

	apiCallWrapper := ApiCallWrapper{
		mimeType: binding.MIMEJSON,
		repo:     repo,
	}

	// init response and context
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	ctx.Params = params

	apiCallWrapper.callApi = func(req any, res any) (*httptest.ResponseRecorder, error) {
		// if req exists, set request params to context
		if req != nil {
			requestJson, err := json.Marshal(req)
			if err != nil {
				return response, err
			}
			request, err := http.NewRequest(method, url, bytes.NewReader(requestJson))
			if err != nil {
				return response, err
			}
			request.Header.Add("Content-Type", apiCallWrapper.mimeType)
			ctx.Request = request
		}
		// get api function
		callApi := api(apiCallWrapper.repo, context)
		callApi(ctx)
		return response, json.Unmarshal(response.Body.Bytes(), res)
	}

	return &apiCallWrapper
}

func TestGetPasswords(t *testing.T) {
	apiWrapper := createApiWrapper(GetPasswords, "GET", "/api/passwords", gin.Params{})

	passwords := []model.Password{}
	_, err := apiWrapper.CallApi(nil, &passwords)
	if err != nil {
		t.Errorf(fmt.Sprintf("failed to call api: %v", err))
	}

	expectedLen := 9
	if len(passwords) != expectedLen {
		t.Errorf("got: %v, expected: %v", len(passwords), expectedLen)
	}
}
