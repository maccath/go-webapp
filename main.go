package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", HelloUser)
	r.HandleFunc("/", HelloWorld)

	http.ListenAndServe(os.Getenv("PORT"), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func HelloUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Hello, %s!", vars["name"])
}
