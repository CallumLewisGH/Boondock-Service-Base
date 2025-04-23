package database

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database wraps the SQL connection pool
type Database struct {
	DBPool *sql.DB
}

var (
	dbInstance *Database
	once       sync.Once // Ensures initialization happens only once
)

// GetDB returns the singleton Database instance
func GetDB() *Database {
	once.Do(func() {
		dbInstance = &Database{}
		dbInstance.InitialiseDatabase()
	})
	return dbInstance
}

func (db *Database) InitialiseDatabase() {
	log.Printf("Connecting to Database...")
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	sqlInfo := os.Getenv("DATABASE_CONNECTION_STRING")
	databaseDriver := os.Getenv("DATABASE_DRIVER")

	pool, err := sql.Open(databaseDriver, sqlInfo)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Configure connection pool
	pool.SetMaxOpenConns(25)
	pool.SetMaxIdleConns(25)
	pool.SetConnMaxLifetime(5 * time.Minute)

	if err = pool.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	db.DBPool = pool
	log.Printf("Database Connection Succeeded")
}

func (db *Database) GetDBPool() *sql.DB {
	return db.DBPool
}
