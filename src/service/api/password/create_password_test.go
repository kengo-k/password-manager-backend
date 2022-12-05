package password

import (
	"testing"

	"github.com/kengo-k/password-manager/model"
)

func TestCreatePassword(t *testing.T) {
	callApi := getGinHandler(CreatePassword)
	createRequest := model.PasswordCreateRequest{
		Name:       "nameX",
		Desc:       "descX",
		CategoryID: "cat2",
	}

	password := model.Password{}
	callApi("POST", "/api/passwords", createRequest, &password)

	expectedName := "nameX"
	if password.Name != expectedName {
		t.Errorf("got: %v, expected: %v", password.Name, expectedName)
	}
}
