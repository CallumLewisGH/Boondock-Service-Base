package query

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	repository "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/repositories"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func MapUserRows(rows *sql.Rows) ([]models.User, error) {
	//Map Results
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.ProfilePicture); err != nil {
			// Log but continue processing other rows
			log.Printf("Failed to scan user row: %v", err)
			continue
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after row mapping iteration: %w", err)
	}

	return users, nil
}

func MapUserRow(row *sql.Row) (*models.User, error) {
	//Map Results
	var user models.User
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.ProfilePicture)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("error not found: %w", err)
	}

	return &user, nil
}

// GetUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	//Create QueryFunction
	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) ([]models.User, error) {
		query, args, err := sb.
			Select("id", "user_name", "email", "profile_picture").
			From("base.users").
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		//Execute Query
		rows, err := conn.QueryContext(ctx, query, args...)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return []models.User{}, nil
			}
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}
		defer rows.Close()

		users, err := MapUserRows(rows)

		return users, err
	}

	return repository.Query(queryFunction)

}

// GetUsers retrieves all users from the database that have the given id
func GetUserById(id uuid.UUID) (*models.User, error) {
	//Create QueryFunction
	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		query, args, err := sb.
			Select("id", "user_name", "email", "profile_picture").
			From("base.users").
			Where(sq.Eq{"id": id}).
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		//Execute Query
		row := conn.QueryRowContext(ctx, query, args...)

		user, err := MapUserRow(row)

		return user, err
	}
	return repository.QueryOne(queryFunction)
}

// GetUsers retrieves all users from the database that have the given id
func GetUserByUserName(userName string) (*models.User, error) {
	//Create QueryFunction
	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		query, args, err := sb.
			Select("id", "user_name", "email", "profile_picture").
			From("base.users").
			Where(sq.Eq{"user_name": userName}).
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		//Execute Query
		row := conn.QueryRowContext(ctx, query, args...)

		user, err := MapUserRow(row)

		return user, err
	}

	return repository.QueryOne(queryFunction)

}
