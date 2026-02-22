package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"x-service/internal/users/usecases"

	myerr "x-service/internal/core/errors"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandlers struct {
	service usecases.Service
}

func NewUserHandlers(userService usecases.Service) *UserHandlers {
	return &UserHandlers{
		service: userService,
	}
}

func (h *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	user, err := toUser(userDTO)
	if err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.service.Add(r.Context(), user); err != nil {
		if errors.Is(err, myerr.ErrUserAlreadyExists) {
			h.respondWithError(w, err, http.StatusConflict)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	userDTO = toResponse(user)
	h.respondWithJSON(w, userDTO, http.StatusOK)
}

func (h *UserHandlers) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.respondWithError(w, myerr.ErrUserNotFound, http.StatusBadRequest)
		return
	}

	user, err := h.service.Get(r.Context(), id)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusNotFound)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	userDTO := toResponse(user)
	h.respondWithJSON(w, userDTO, http.StatusOK)
}

func (h *UserHandlers) HandleUpdateUserPassword(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.respondWithError(w, myerr.ErrUserNotFound, http.StatusBadRequest)
		return
	}

	var passwordUserDTO PasswordUserDTO
	if err := json.NewDecoder(r.Body).Decode(&passwordUserDTO); err != nil {
		h.respondWithError(w, err, http.StatusBadRequest)
		return
	}

	if err := h.service.UpdatePassword(r.Context(), id, passwordUserDTO.GetContent()); err != nil {
		switch {
		case errors.Is(err, myerr.ErrUserNotFound):
			h.respondWithError(w, err, http.StatusNotFound)
		case errors.Is(err, myerr.ErrPasswordIsEmpty),
			errors.Is(err, myerr.ErrPasswordTooShort):
			h.respondWithError(w, err, http.StatusBadRequest)
		default:
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	user, err := h.service.Get(r.Context(), id)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusBadRequest)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	userDTO := toResponse(user)
	h.respondWithJSON(w, userDTO, http.StatusOK)
}

func (h *UserHandlers) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.respondWithError(w, myerr.ErrUserNotFound, http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			h.respondWithError(w, err, http.StatusNotFound)
		} else {
			h.respondWithError(w, err, http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandlers) respondWithError(w http.ResponseWriter, err error, statusCode int) {
	errorDTO := NewErrorDTO(err.Error())
	http.Error(w, errorDTO.ToString(), statusCode)
}

func (*UserHandlers) respondWithJSON(w http.ResponseWriter, userDTO UserDTO, statusCode int) {
	w.WriteHeader(statusCode)
	if _, err := w.Write(userDTO.ToBytes()); err != nil {
		fmt.Println("failed to write http-response: ", err)
		return
	}
}
