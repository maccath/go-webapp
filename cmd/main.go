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
	loadEnv()

	var repo = memory.Repository{}
	repo.Save(memory.User{Name: "katy"})

	var g = greeting.NewService(repo)

	http.ListenAndServe(":"+getPort(), web.Handler(g))
}

func getPort() string {
	port, specified := os.LookupEnv("PORT")
	if !specified {
		port = "80"
	}
	return port
}

func loadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
