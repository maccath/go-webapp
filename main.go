package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", HelloUser)
	r.HandleFunc("/", HelloWorld)

	http.ListenAndServe(":80", r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func HelloUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Hello, %s!", vars["name"])
}
