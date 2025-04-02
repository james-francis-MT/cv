package handlers

import (
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Contact("a@b.com")
	err := templates.Layout(c, "contact").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
