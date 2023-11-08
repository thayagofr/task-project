package rest

import (
	"encoding/json"
	"net/http"
	"task-api/internal/adapters/in/rest/request"
	"task-api/internal/adapters/in/rest/response"
	"task-api/internal/ports/in"
)

type AuthHandler struct {
	registerUseCase in.RegisterUser
}

func NewAuthHandler(registerUseCase in.RegisterUser) *AuthHandler {
	return &AuthHandler{registerUseCase: registerUseCase}
}

func (handler *AuthHandler) Register(writer http.ResponseWriter, req *http.Request) {
	var newUser request.NewUser

	if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
		respondWithJSON(writer, response.NewError(req.RequestURI, "Invalid payload format"), http.StatusBadRequest)
		return
	}

	userToCreate := newUser.ToDomain()
	createdUser, err := handler.registerUseCase.Register(req.Context(), userToCreate)
	if err != nil {
		respondWithJSON(writer, response.NewError(req.RequestURI, err.Error()), http.StatusBadRequest)
		return
	}

	respondWithJSON(writer, response.FromUserDomain(createdUser), http.StatusCreated)
}

func (handler *AuthHandler) Authenticate(writer http.ResponseWriter, req *http.Request) {
	var credentials request.LoginCredentials

	if err := json.NewDecoder(req.Body).Decode(&credentials); err != nil {
		respondWithJSON(writer, response.NewError(req.RequestURI, "Invalid payload format"), http.StatusBadRequest)
		return
	}

	accessCredentials, err := handler.registerUseCase.Authenticate(req.Context(), credentials.ToDomain())
	if err != nil {
		respondWithJSON(writer, response.NewError(req.RequestURI, err.Error()), http.StatusBadRequest)
		return
	}

	respondWithJSON(writer, response.FromAccessDomain(accessCredentials), http.StatusOK)
}
