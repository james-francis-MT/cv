package handlers

import (
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Layout()
	err := c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
