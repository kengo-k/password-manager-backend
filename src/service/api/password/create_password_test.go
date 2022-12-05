package password

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/kengo-k/password-manager/model"
)

func TestCreatePassword(t *testing.T) {
	callApi, response, ctx := getGinHandler(CreatePassword)
	createRequest := model.PasswordCreateRequest{
		Name:       "nameX",
		Desc:       "descX",
		CategoryID: "cat2",
	}
	createRequestJson, err := json.Marshal(&createRequest)
	if err != nil {
		fmt.Printf("marshal?: %v\n", err)
	}
	req, _ := http.NewRequest("POST", "/api/passwords", bytes.NewReader(createRequestJson))
	req.Header.Add("Content-Type", binding.MIMEJSON)
	ctx.Request = req
	callApi(ctx)

	password := model.Password{}
	if err := json.Unmarshal(response.Body.Bytes(), &password); err != nil {
		t.Errorf("failed to decode json")
	}

	expectedName := "nameX"
	if password.Name != expectedName {
		t.Errorf("got: %v, expected: %v", password.Name, expectedName)
	}
}
