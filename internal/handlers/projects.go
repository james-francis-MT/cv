package handlers

import (
	"github.com/james-francis-MT/cv/internal/models"
	"github.com/james-francis-MT/cv/internal/templates"
	"net/http"
)

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := []models.Project{
		{
			Title:       "CV Website",
			Description: "A website to show off my skills!",
			TechStack:   []string{"Go", "TailwindCSS", "Templ", "HTMX"},
			Link:        "https://github.com/james-francis-MT/cv",
		},
		{
			Title:       "Building my own shell",
			Description: "Following a guide from CodeCrafters to build my own shell in Kotlin",
			TechStack:   []string{"Kotlin"},
			Link:        "https://github.com/james-francis-MT/kotlin-shell",
		},
	}

	c := templates.Projects(projects)
	err := templates.Layout(c, "projects").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
