package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"x-service/internal/users/usecases"

	myerr "x-service/internal/core/errors"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	service usecases.Service
}

func NewHTTPHandlers(userService *usecases.UserService) *HTTPHandlers {
	return &HTTPHandlers{
		service: userService,
	}
}

func (h *HTTPHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	user, err := ToUser(userDTO)
	if err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.service.Add(user); err != nil {
		if errors.Is(err, myerr.ErrUserAlreadyExists) {
			h.respondWithError(w, err, http.StatusConflict)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	userDTO = toResponse(user)
	h.respondWithjSON(w, userDTO, http.StatusOK)
}

func (h *HTTPHandlers) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	user, err := h.service.Get(username)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusBadRequest)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	userDTO := toResponse(user)
	h.respondWithjSON(w, userDTO, http.StatusOK)
}

func (h *HTTPHandlers) HandleUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	var passwordUserDTO PasswordUserDTO
	if err := json.NewDecoder(r.Body).Decode(&passwordUserDTO); err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	username := mux.Vars(r)["username"]
	if err := h.service.UpdatePassword(username, passwordUserDTO.GetContent()); err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusBadRequest)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	user, err := h.service.Get(username)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusBadRequest)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)

		}

		return
	}

	userDTO := toResponse(user)
	h.respondWithjSON(w, userDTO, http.StatusOK)
}

func (h *HTTPHandlers) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	if err := h.service.Delete(username); err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusBadRequest)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) respondWithError(w http.ResponseWriter, err error, statusCode int) {
	errorDTO := NewErrorDTO(err.Error())
	http.Error(w, errorDTO.ToString(), statusCode)
}

func (h *HTTPHandlers) respondWithjSON(w http.ResponseWriter, userDTO UserDTO, statusCode int) {
	w.WriteHeader(statusCode)
	if _, err := w.Write(userDTO.ToBytes()); err != nil {
		fmt.Println("failed to write http-response: ", err)
		return
	}
}
