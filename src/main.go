package main

import (
	"github.com/kengo-k/password-manager/server"
)

func main() {
	service := server.NewService()
	server := server.NewServer(service)
	server.Run(":8080")
}
