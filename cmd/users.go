package main

import (
	"fmt"
	"x-service/internal/users/adapters"
	"x-service/internal/users/handlers"
	"x-service/internal/users/usecases"
)

func main() {
	repository := adapters.NewStorage()
	service := usecases.NewUserService(repository)
	httpHandlers := handlers.NewHTTPHandlers(service)
	server := handlers.NewHTTPServer(httpHandlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("failed to start http-server")
	}
}
