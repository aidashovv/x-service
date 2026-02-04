package main

import (
	"fmt"
	"pdd/internal/users/adapters"
	"pdd/internal/users/handlers/http"
	"pdd/internal/users/services"
)

func main() {
	repository := adapters.NewStorage()
	service := services.NewUserService(repository)
	handlers := http.NewHTTPHandlers(service)
	server := http.NewHTTPServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("failed to start http-server")
	}
}
