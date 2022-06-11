package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maccath/go-webapp/repository"
	"github.com/maccath/go-webapp/users"
	"net/http"
	"os"
)

var repo users.Repository = &repository.UserMemoryRepo{}

func init() {
	repo.Save(users.User{Name: "katy"})
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
	var user *users.User

	if vars["name"] != "" {
		user = repo.FindByName(vars["name"])
	}

	fmt.Fprintf(w, Hello(user))
}

func Hello(user *users.User) string {
	if user == nil {
		return "Hello, world!"
	} else {
		return "Hello, " + user.Name + "!"
	}
}
