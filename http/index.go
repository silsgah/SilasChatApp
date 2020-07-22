package main

import (
	"net/http"

	"github.com/mattn/go/src/fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you requested : %\n", r.URL.Path)
	})

	http.ListenAndServe(":90", nil)
}
