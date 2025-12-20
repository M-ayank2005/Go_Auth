package main

import (
	"log"
	"net/http"
	"os"

	"GO_Auth/db"
	"GO_Auth/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set")
	}



	db.Connect()

	handlers := routes.RegisterRoutes()

	log.Println("Server running on Port :8080")
	if err := http.ListenAndServe(":8080", handlers); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}

