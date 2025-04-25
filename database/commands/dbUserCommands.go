package dbCommand

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CallumLewisGH/Boondock-Service-Base/database/mappers"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	repository "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/repositories"
	sq "github.com/Masterminds/squirrel"
)

// Creates a new user in the
func CreateUser(user models.UserDbRequest) (*models.User, error) {
	// Create QueryFunction
	commandFunction := func(tx *sql.Tx, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		query, args, err := sb.
			Insert("base.users").
			Columns("user_name", "email", "password_hash").
			Values(user.UserName, user.Email, user.PasswordHash).
			Suffix("RETURNING id, created_utc, user_name, email, profile_picture").
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		// Execute Query
		row := tx.QueryRowContext(ctx, query, args...)

		if row.Err() != nil {
			return nil, fmt.Errorf("failed to execute query: %w", row.Err())
		}

		user, err := mappers.MapUserRow(row)

		if err != nil {
			return nil, fmt.Errorf("failed to scan query: %w", err)
		}

		return user, nil
	}

	return repository.ExecuteOne(commandFunction)
}

func UpdateUserById(model models.UserDbRequest) (*models.User, error) {
	commandFunction := func(tx *sql.Tx, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		// Start with base update
		update := sb.Update("base.users")

		fmt.Printf("Before DB: %+v\n", model)

		// Conditionally add each field
		if model.UserName != nil {
			update = update.Set("user_name", model.UserName)
		}
		if model.Email != nil {
			update = update.Set("email", model.Email)
		}
		if model.PasswordHash != nil {
			update = update.Set("password_hash", model.PasswordHash)
		}
		if model.ProfilePicture != nil {
			update = update.Set("profile_picture", model.ProfilePicture)
		}
		if model.DeletedUTC != nil {
			update = update.Set("deleted_utc", model.DeletedUTC)
		}

		// Complete the query
		query, args, err := update.
			Where(sq.Eq{"id": model.ID}).
			Suffix("RETURNING id, created_utc, user_name, email, profile_picture").
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		// Execute Query
		row := tx.QueryRowContext(ctx, query, args...)

		if row.Err() != nil {
			return nil, fmt.Errorf("failed to execute query: %w", row.Err())
		}

		user, err := mappers.MapUserRow(row)
		if err != nil {
			return nil, fmt.Errorf("failed to scan query: %w", err)
		}

		return user, nil
	}

	return repository.ExecuteOne(commandFunction)
}
