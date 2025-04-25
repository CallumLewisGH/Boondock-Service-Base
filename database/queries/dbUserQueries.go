package dbQuery

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/CallumLewisGH/Boondock-Service-Base/database/mappers"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	repository "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/repositories"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// GetUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	//Create QueryFunction
	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) ([]models.User, error) {
		query, args, err := sb.
			Select("id", "created_utc", "user_name", "email", "profile_picture").
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

		users, err := mappers.MapUserRows(rows)

		return users, err
	}

	return repository.Query(queryFunction)

}

// GetUsers retrieves all users from the database that have the given id
func GetUserById(id uuid.UUID) (*models.User, error) {
	//Create QueryFunction
	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		query, args, err := sb.
			Select("id", "created_utc", "user_name", "email", "profile_picture").
			From("base.users").
			Where(sq.Eq{"id": id}).
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		//Execute Query
		row := conn.QueryRowContext(ctx, query, args...)

		user, err := mappers.MapUserRow(row)

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

		user, err := mappers.MapUserRow(row)

		return user, err
	}

	return repository.QueryOne(queryFunction)

}
