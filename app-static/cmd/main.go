package main

import (
	"app-static/api"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	godotenv.Load()

	server := api.ApiServer{
		Addr: ":" + os.Getenv("PORT_APP_STATIC"),
	}

	return server.Run()
}
