package handlers

import (
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Projects("a@b.com")
	err := templates.Layout(c, "projects").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
