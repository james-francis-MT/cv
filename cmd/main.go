package main

import (
	"fmt"
	"github.com/james-francis-MT/cv/internal/handlers"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there")
}

func projectHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there")
}

func contactHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/*", fileServer)

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/project", projectHandler)
	http.HandleFunc("/contact", contactHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
