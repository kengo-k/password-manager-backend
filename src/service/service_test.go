package service_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	passwordContext "github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/model"
	"github.com/kengo-k/password-manager/server"
	"github.com/kengo-k/password-manager/service"
)

func setup() *http.Server {
	config := env.NewConfig("testdata/.test.env")
	context := passwordContext.NewContext(runmode.FILE_TO_FILE, config)
	service := service.NewServiceProvider(context)
	router := server.NewServer(service)
	server := &http.Server{Addr: ":8080", Handler: router}
	//server.ListenAndServe()
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	server := setup()
	go func() {
		server.ListenAndServe()
	}()
	m.Run()
	//server.Shutdown(context.Background())
	//os.Exit(code)
}

func TestGetPasswords(t *testing.T) {
	client := resty.New()
	var results []*model.Password = []*model.Password{}
	_, err := client.R().
		EnableTrace().
		SetResult(&results).
		Get("http://localhost:8080/api/passwords")

	if err != nil {
		t.Errorf("failed to call get api")
	}

	expectedLen := 9
	if len(results) != expectedLen {
		t.Errorf("got: %v, expected: %v", len(results), expectedLen)
	}

	names := []string{
		"name1", "name2", "name3", "name10", "name20", "name30", "nameA", "nameB", "nameC",
	}

	for i, password := range results {
		if password.Name != names[i] {
			t.Errorf("got: %v, expected: %v", password.Name, names[i])
		}
	}

}
