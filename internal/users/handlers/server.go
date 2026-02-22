package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(port string, httpHandlers *UserHandlers) *HTTPServer {
	router := mux.NewRouter()

	router.Path("/users").Methods("POST").HandlerFunc(httpHandlers.HandleCreateUser)
	router.Path("/users/{id}").Methods("GET").HandlerFunc(httpHandlers.HandleGetUser)
	router.Path("/users/{id}").Methods("PATCH").HandlerFunc(httpHandlers.HandleUpdateUserPassword)
	router.Path("/users/{id}").Methods("DELETE").HandlerFunc(httpHandlers.HandleDeleteUser)

	return &HTTPServer{
		server: &http.Server{
			Addr:           ":" + port,
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
