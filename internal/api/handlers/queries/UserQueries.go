package query

import (
	"fmt"

	dbQuery "github.com/CallumLewisGH/Boondock-Service-Base/database/queries"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	"github.com/google/uuid"
)

// GetUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	users, err := dbQuery.GetAllUsers()

	if err != nil {
		return nil, fmt.Errorf("failed to call query: %w", err)
	}

	return users, nil

}

// GetUsers retrieves all users from the database that have the given id
func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := dbQuery.GetUserById(id)

	if err != nil {
		return nil, fmt.Errorf("failed to call query: %w", err)
	}

	return user, nil
}

// GetUsers retrieves all users from the database that have the given id
func GetUserByUserName(userName string) (*models.User, error) {
	user, err := dbQuery.GetUserByUserName(userName)

	if err != nil {
		return nil, fmt.Errorf("failed to call query: %w", err)
	}

	return user, nil
}
