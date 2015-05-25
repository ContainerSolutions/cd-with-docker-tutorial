package main

import (
	"fmt"
	"net/http"
)

func homeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
}

func main() {
	http.Handle("/", homeHandler())
	http.ListenAndServe(":80", nil)
}
