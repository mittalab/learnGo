package main

import (
		"net/http"
	"fmt"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func about_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}