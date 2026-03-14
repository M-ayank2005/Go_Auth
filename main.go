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
		log.Println("No .env file found, using system environment variables")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	db.Connect()

	handlers := routes.RegisterRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on Port :%s\n", port)
	if err := http.ListenAndServe(":"+port, handlers); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

