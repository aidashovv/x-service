package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/users").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateUser)
	router.Path("/users/{username}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetUser)
	router.Path("/users/{username}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandleUpdateUserPassword)
	router.Path("/users/{username}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandleDeleteUser)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil // nil, так как ErrServerClosed - базовая ошибка, когда все четко
		}

		return err
	}

	return nil
}
