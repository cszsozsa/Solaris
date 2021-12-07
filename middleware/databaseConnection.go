package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Creating the database connection
func createConnection() *sql.DB {
	// Loading env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Opening the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// Checking the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connection successfull!")
	// Return the connection
	return db
}
