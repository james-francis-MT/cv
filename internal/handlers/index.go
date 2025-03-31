package handlers

import (
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	c := templates.Index("a@b.com")
	c.Render(r.Context(), w)
}
