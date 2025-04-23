package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DBPool *sql.DB
)

func InitialiseDatabase() {
	log.Printf("Connecting to Database...")
	err := godotenv.Load()

	if err != nil {
		log.Printf("Error loading .env file")
	}

	var sqlInfo string = os.Getenv("DATABASE_CONNECTION_STRING")
	var databaseDriver string = os.Getenv("DATABASE_DRIVER")

	db, err := sql.Open(databaseDriver, sqlInfo)
	if err != nil {
		panic(err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	DBPool = db

	log.Printf("Database Connection Succeeded")
}
