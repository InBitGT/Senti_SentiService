package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"SentiService/db"
	"SentiService/internal/config"
	"SentiService/internal/modules/address"
	"SentiService/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	config.Init()

	database := db.Database()

	// AutoMigrate addresses
	if err := database.AutoMigrate(&address.Address{}); err != nil {
		log.Fatalf("Error en migracion address: %v", err)
	}

	srv := server.NewServer(database)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	serverHTTP := &http.Server{
		Addr:         ":" + port,
		Handler:      srv.Router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("SentiService escuchando en puerto %s", port)
	log.Fatal(serverHTTP.ListenAndServe())
}
