package handlers

import (
	"encoding/json"
	"net/http"

	"fakorede-bolu/go-ach/cmd/todo"
	"fakorede-bolu/go-ach/pkg/helpers"
	"fakorede-bolu/go-ach/pkg/middlewares"

	"github.com/go-chi/chi"
)

// Handler defines interface for todo
type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

// TodoHandler struct
type TodoHandler struct {
	Service todo.Service
}

func (h *TodoHandler) Router() chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.SecureHeaders)
	r.Use(middlewares.LogRequest)
	r.Use(middlewares.RecoverPanic)

	r.Route("/todo", func(r chi.Router) {
		r.Get("/", h.Get)
		r.Post("/", h.Create)
		r.Get("/{id}", h.GetByID)

	})

	return r
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	todos, err := h.Service.FindAllTodos()

	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := h.Service.FindTodoByID(id)

	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	var t todo.Model

	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	todo, err := h.Service.CreateTodo(&t)

	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo)

}
