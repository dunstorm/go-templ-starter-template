package handlers

import (
	"net/http"

	"github.com/dunstorm/go-templ-starter-template/internal/templates"
)

type GetHomeHandler struct{}

func NewHomeHandler() *GetHomeHandler {
	return &GetHomeHandler{}
}

func (gh *GetHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Index()
	err := templates.Layout(c, "Home").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
	}
}
