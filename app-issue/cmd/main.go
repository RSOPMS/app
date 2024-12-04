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

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	dbDataSource := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s", dbName, dbUser, dbPassword, dbHost, dbPort, dbSslMode)
	db, err := sql.Open("postgres", dbDataSource)
	if err != nil {
		return err
	}

	server := api.ApiServer{
		Addr: ":" + os.Getenv("PORT_APP_ISSUE"),
		Db:   db,
	}

	return server.Run()
}
