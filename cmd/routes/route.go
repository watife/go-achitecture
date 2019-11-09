package routes

import (
	"fakorede-bolu/go-ach/cmd/handlers"
	"fakorede-bolu/go-ach/cmd/todo"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	Todo todo.Service
	// Logger kitlog.Logger

	router chi.Router
}

// NewRouter returns a new router for the api
func NewRouter(td todo.Service) *Server {
	s := &Server{
		Todo: td,
	}

	r := chi.NewRouter()

	r.Use(accessControl)

	r.Route("/go-ach", func(r chi.Router) {
		h := handlers.TodoHandler{Service: s.Todo}
		r.Mount("/v1", h.Router())
	})

	s.router = r

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
