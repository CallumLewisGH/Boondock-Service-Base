package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/CallumLewisGH/Boondock-Service-Base/database"
)

var (
	dbPool = database.DBPool
)

func Query[T any](f func(*sql.Conn, context.Context) ([]T, error)) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := dbPool.Conn(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Warning: failed to close database connection: %v", err)
		}
	}()

	result, err := f(conn, ctx)

	return result, err

}

// WithTransaction executes a function within a transaction
func Command[T any](f func(*sql.Tx, context.Context) ([]T, error)) ([]T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := dbPool.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // re-throw panic after rollback
		}
	}()

	result, err := f(tx, ctx)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return nil, fmt.Errorf("operation error: %v, rollback error: %w", err, rbErr)
		}
		return nil, err
	}

	tx.Commit()

	return result, err
}
