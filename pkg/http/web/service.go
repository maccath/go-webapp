package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maccath/go-webapp/pkg/greeting"
	"net/http"
)

func Handler(g greeting.Service) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", greetUser(g))
	r.HandleFunc("/", greetUser(g))

	return r
}

func greetUser(g greeting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, g.Greet(mux.Vars(r)["name"]))
	}
}
