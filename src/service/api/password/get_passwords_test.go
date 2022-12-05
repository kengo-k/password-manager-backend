package password

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
	"github.com/kengo-k/password-manager/types"
)

// get api function for test
func getGinHandler(api types.ApiCall) (gin.HandlerFunc, *httptest.ResponseRecorder, *gin.Context) {
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
	handler := api(repo, context)
	// init response and context
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	return handler, response, ctx
}

func TestGetPasswords(t *testing.T) {
	callApi, response, ctx := getGinHandler(GetPasswords)
	callApi(ctx)
	passwords := []model.Password{}
	if err := json.Unmarshal(response.Body.Bytes(), &passwords); err != nil {
		t.Errorf("failed to decode json")
	}
	expectedLen := 9
	if len(passwords) != expectedLen {
		t.Errorf("got: %v, expected: %v", len(passwords), expectedLen)
	}
}
