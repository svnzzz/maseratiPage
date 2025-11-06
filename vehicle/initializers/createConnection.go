package initializers

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func CreateConnection() error {
	connectionString := LoadEnvVariables()

	DB, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return fmt.Errorf("error while opening the connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error while pinging the database: %v", err)
	}

	fmt.Print("Successfuly connected to the database")

	return nil
}
