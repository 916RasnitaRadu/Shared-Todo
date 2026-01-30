package app

import (
	"net/http"
	"shared-todo/view"

	"github.com/go-chi/chi/v5"
)

func HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	view.HelloWorld("World").Render(r.Context(), w)
}

func HandleHelloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	view.HelloWorld(name).Render(r.Context(), w)
}
