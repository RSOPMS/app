package main

import (
	"app-issue/api"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	godotenv.Load()

	db, err := sql.Open("postgres", "dbname=bugbase user=bugbase password=password host=localhost port=5432 sslmode=disable")
	if err != nil {
		return err
	}

	server := api.ApiServer{
		Addr: ":" + os.Getenv("APP_PORT"),
		Db:   db,
	}

	return server.Run()
}
