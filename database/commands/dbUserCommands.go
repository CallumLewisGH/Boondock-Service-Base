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
func CreateUser(user models.User) (*models.User, error) {
	// Create QueryFunction
	commandFunction := func(tx *sql.Tx, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
		query, args, err := sb.
			Insert("base.users").
			Columns("id", "user_name", "email", "hash_salt", "password_hash", "profile_picture").
			Values(user.ID, user.UserName, user.Email, user.PasswordHash, user.ProfilePicture, user.CreatedUTC).
			Suffix("RETURNING id, user_name, email, profile_picture").
			ToSql()

		if err != nil {
			return nil, fmt.Errorf("failed to build query: %w", err)
		}

		// Execute Query
		row := tx.QueryRowContext(ctx, query, args...)

		user, err := mappers.MapUserRow(row)

		if err != nil {
			return nil, fmt.Errorf("failed to execute query: %w", err)
		}

		return user, nil
	}

	return repository.ExecuteOne(commandFunction)
}

// // GetUsers retrieves all users from the database that have the given id
// func GetUserById(id uuid.UUID) (*models.User, error) {
// 	//Create QueryFunction
// 	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
// 		query, args, err := sb.
// 			Select("id", "user_name", "email", "profile_picture").
// 			From("base.users").
// 			Where(sq.Eq{"id": id}).
// 			ToSql()

// 		if err != nil {
// 			return nil, fmt.Errorf("failed to build query: %w", err)
// 		}

// 		//Execute Query
// 		row := conn.QueryRowContext(ctx, query, args...)

// 		user, err := MapUserRow(row)

// 		return user, err
// 	}
// 	return repository.QueryOne(queryFunction)
// }

// // GetUsers retrieves all users from the database that have the given id
// func GetUserByUserName(userName string) (*models.User, error) {
// 	//Create QueryFunction
// 	queryFunction := func(conn *sql.Conn, ctx context.Context, sb sq.StatementBuilderType) (*models.User, error) {
// 		query, args, err := sb.
// 			Select("id", "user_name", "email", "profile_picture").
// 			From("base.users").
// 			Where(sq.Eq{"user_name": userName}).
// 			ToSql()

// 		if err != nil {
// 			return nil, fmt.Errorf("failed to build query: %w", err)
// 		}

// 		//Execute Query
// 		row := conn.QueryRowContext(ctx, query, args...)

// 		user, err := MapUserRow(row)

// 		return user, err
// 	}

// 	return repository.QueryOne(queryFunction)

// }
