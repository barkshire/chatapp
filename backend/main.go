package main

import (
	"fmt"
	"net/http"
)

func router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server up")
	})
}

func main() {
	router()
	http.ListenAndServe(":8080", nil)
}
