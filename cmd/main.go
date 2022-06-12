package main

import (
	"github.com/joho/godotenv"
	"github.com/maccath/go-webapp/pkg/greeting"
	"github.com/maccath/go-webapp/pkg/http/web"
	"github.com/maccath/go-webapp/pkg/storage/memory"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var repo = memory.Repository{}
	repo.Save(memory.User{Name: "katy"})

	var g = greeting.NewService(repo)

	port := os.Getenv("PORT")

	http.ListenAndServe(":"+port, web.Handler(g))
}
