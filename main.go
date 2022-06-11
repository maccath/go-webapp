package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", SayHello)
	r.HandleFunc("/", SayHello)

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, r)
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, Hello(vars["name"]))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return "Hello, " + name + "!"
}
