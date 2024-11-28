package main

import (
	"app-issue/api"
	"database/sql"
	"fmt"
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

	dbName := "bugbase"
	dbUser := "bugbase"
	dbPassword := "password"
	dbHost := os.Getenv("DB_HOST")
	dbPort := 5432
	dbSslMode := "disable"

	dbDataSource := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s", dbName, dbUser, dbPassword, dbHost, dbPort, dbSslMode)
	db, err := sql.Open("postgres", dbDataSource)
	if err != nil {
		return err
	}

	server := api.ApiServer{
		Addr: ":" + os.Getenv("APP_PORT"),
		Db:   db,
	}

	return server.Run()
}
