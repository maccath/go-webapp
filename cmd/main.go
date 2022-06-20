package main

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/maccath/go-webapp/pkg/greeting"
	"github.com/maccath/go-webapp/pkg/http/web"
	sqlStorage "github.com/maccath/go-webapp/pkg/storage/sql"
	"log"
	"net/http"
	"os"
)

func main() {
	loadEnv()

	connStr := os.Getenv("PGSQL")

	// Connect to database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	var repo = sqlStorage.NewRepository(db)

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
