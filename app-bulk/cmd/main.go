package main

import (
	"app-bulk/api"
	"app-bulk/pkg"
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

	server := api.ApiServer{
		Addr: ":" + os.Getenv("PORT_APP_BULK"),
	}

	// Initialize NATS connection
	err := pkg.InitNATS()
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer pkg.CloseNATSConnection()

	return server.Run()
}
