package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todolist-api/internal/models"
	"todolist-api/internal/services"
	"todolist-api/pkg/utils"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	registeredUser, err := h.service.CreateUser(r.Context(), &user)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := utils.GenerateAccessToken(registeredUser.ID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate access token: %v", err))
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(registeredUser.ID)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to generate refresh token: %v", err))
		return
	}

	sendJSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    map[string]interface{}{
			"user": registeredUser,
			"access_token": accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

}