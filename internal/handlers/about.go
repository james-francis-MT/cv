package handlers

import (
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.About()
	err := templates.Layout(c, "about").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
