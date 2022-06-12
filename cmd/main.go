package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maccath/go-webapp/pkg/storage/memory"
	"github.com/maccath/go-webapp/pkg/user"
	"net/http"
	"os"
)

var repo user.Repository = &memory.Repository{}

func init() {
	repo.Save(user.Model{Name: "katy"})
}

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
	var user *user.Model

	if vars["name"] != "" {
		user = repo.FindByName(vars["name"])
	}

	fmt.Fprintf(w, Hello(user))
}

func Hello(user *user.Model) string {
	if user == nil {
		return "Hello, world!"
	} else {
		return "Hello, " + user.Name + "!"
	}
}
