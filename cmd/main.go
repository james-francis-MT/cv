package main

import (
	"fmt"
	"github.com/james-francis-MT/cv/internal/handlers"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/{$}", handlers.IndexHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/projects", handlers.ProjectsHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
