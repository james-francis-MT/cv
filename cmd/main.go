package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello there")
}

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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/project", projectHandler)
	http.HandleFunc("/contact", contactHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
