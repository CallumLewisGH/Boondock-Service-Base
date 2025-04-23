package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	"github.com/Masterminds/squirrel"
)

// GetUsers retrieves all users from the database
func GetUsers(conn *sql.Conn, ctx context.Context, sb squirrel.StatementBuilderType) ([]models.User, error) {
	// Build and execute query
	query, args, err := sb.
		Select("id", "user_name", "email").
		From("base.users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.User{}, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email); err != nil {
			// Log but continue processing other rows
			log.Printf("Failed to scan user row: %v", err)
			continue
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row iteration: %w", err)
	}

	return users, nil
}
