package main

import (
	"app-issue/api"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	godotenv.Load()

	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		return err
	}

	server := api.ApiServer{
		Addr: ":" + os.Getenv("APP_PORT"),
		Db:   db,
	}

	return server.Run()
}
