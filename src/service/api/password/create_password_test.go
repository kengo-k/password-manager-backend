package password

import (
	"fmt"
	"testing"

	"github.com/kengo-k/password-manager/model"
)

func TestCreatePassword(t *testing.T) {
	callApi := createApiWrapper(CreatePassword)

	request := model.PasswordCreateRequest{
		Name:       "nameX",
		Desc:       "descX",
		CategoryID: "cat2",
	}

	password := model.Password{}
	err := callApi("POST", "/api/passwords", request, &password)
	if err != nil {
		t.Errorf(fmt.Sprintf("failed to call api: %v", err))
	}

	expectedName := "nameX"
	if password.Name != expectedName {
		t.Errorf("got: %v, expected: %v", password.Name, expectedName)
	}
}
