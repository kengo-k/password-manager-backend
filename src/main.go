package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/server"
	"github.com/kengo-k/password-manager/service"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config, err := env.NewConfig(".env")
	if err != nil {
		panic("failed to get config, abort!")
	}
	context := context.NewContext(runmode.GIT_TO_GIT, config)
	service := service.NewServiceProvider(context)
	router := server.NewServer(service)
	server := &http.Server{Addr: ":8080", Handler: router}
	server.ListenAndServe()
}
