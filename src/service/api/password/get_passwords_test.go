package password

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/repo"
)

func TestGetPasswords(t *testing.T) {

	config := env.NewConfig("testdata/.test.env")

	context := context.NewContext(runmode.FILE_TO_FILE, config)
	passwords, _ := context.Load()

	database := model.NewDatabase()
	database.Init(passwords)

	repo := repo.NewRepository(database)

	callApi := GetPasswords(repo, context)
	response := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(response)
	callApi(ctx)

	fmt.Println(response.Body)
	//buf := bytes.NewBuffer(response.Body.Bytes())
	//dec := gob.NewDecoder(buf)
	passwords2 := []model.Password{}

	if err := json.Unmarshal(response.Body.Bytes(), &passwords2); err != nil {
		log.Fatal(err)
	}

	//err := dec.Decode(&passwords2)
	//fmt.Printf("err? %v\n", err)

	fmt.Printf("body: %v\n", passwords2)
}
