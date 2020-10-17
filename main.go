package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if "/favicon.ico" != r.URL.Path {
			log.Print("called ", r.URL.Path)
		}

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HI~!!!")
	})

	http.Handle("/test", http.FileServer(http.Dir("./views/test")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
