package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() (connectionString string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Errore: %v", err)
	}

	connectionString = os.Getenv("SQL_CONNECTION_STRING")

	if connectionString == "" {
		log.Fatal("connectionString not found")
	}
	return connectionString
}
