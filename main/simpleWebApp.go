package main

import (
		"net/http"
	"fmt"
	"io/ioutil"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
	fmt.Fprintf(w, "You %s add <strong>Variables</strong>", "can")

	fmt.Fprintf(w, `
	<h1>Whoa, Go is neat!</h1>
	You %s add <strong>Variables</strong>
	`, "HELLO")
}

func about_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ABOUT!\n")

	/*
	Hitting a URL and fetching data
	 */
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	resp.Body.Close()

	fmt.Fprintf(w, string_body)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)

	http.ListenAndServe(":8000", nil)
}