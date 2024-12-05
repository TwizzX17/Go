package main

import (
	"log"
	dbSetup "stock_data/internal/db"
	"stock_data/internal/router"

	"github.com/joho/godotenv"

	/* "my-go-api/internal/router" */
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	err = dbSetup.InitDB()
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
		return
	}

	r := router.SymbolsRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
