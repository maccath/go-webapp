package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		fmt.Fprintf(w, "Hello, %s!", vars["name"])
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":80", r)
}
