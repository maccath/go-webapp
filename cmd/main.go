package main

import (
	"github.com/maccath/go-webapp/pkg/greeting"
	"github.com/maccath/go-webapp/pkg/http/web"
	"github.com/maccath/go-webapp/pkg/storage/memory"
	"net/http"
	"os"
)

func main() {
	var repo = memory.Repository{}
	repo.Save(memory.User{Name: "katy"})

	var g = greeting.NewService(repo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	http.ListenAndServe(":"+port, web.Handler(g))
}
