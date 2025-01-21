package handlers

import (
	"encoding/json"
	"net/http"
	"todolist-api/internal/models"
	"todolist-api/internal/services"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	service *services.TodoService
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("sub").(string)

	var input models.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userId

	todo, err := h.service.CreateTodo(r.Context(), &input)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusCreated, Response{
		Success: true,
		Data:    todo,
	})
}

func (h *TodoHandler) GetTodoByUserID(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID").(string)

	todo, err := h.service.GetTodoByUserID(r.Context(), userId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    todo,
	})
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	todo, err := h.service.GetTodoById(r.Context(), todoId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    todo,
	})
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	var input models.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}

	todo, err := h.service.UpdateTodo(r.Context(), todoId, &input)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    todo,
	})
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	err := h.service.DeleteTodo(r.Context(), todoId)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(w, http.StatusOK, Response{
		Success: true,
	})
}

func sendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendError(w http.ResponseWriter, statusCode int, message string) {
	sendJSON(w, statusCode, Response{
		Success: false,
		Error:   message,
	})
}
