package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDatabase() {
	fmt.Println("Connecting to Database...")
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	var sqlInfo string = os.Getenv("DATABASE_CONNECTION_STRING")
	var databaseDriver string = os.Getenv("DATABASE_DRIVER")

	db, err := sql.Open(databaseDriver, sqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database Connection Succeeded")
}
