package main

import (
	"net/http"

	"github.com/kengo-k/password-manager/context"
	"github.com/kengo-k/password-manager/context/runmode"
	"github.com/kengo-k/password-manager/env"
	"github.com/kengo-k/password-manager/server"
	"github.com/kengo-k/password-manager/service"
)

func main() {
	config := env.NewConfig(".env")
	context := context.NewContext(runmode.GIT_TO_GIT, config)
	service := service.NewServiceProvider(context)
	router := server.NewServer(service)
	server := &http.Server{Addr: ":8080", Handler: router}
	server.ListenAndServe()
}
