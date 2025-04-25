package mappers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
)

func MapUserRows(rows *sql.Rows) ([]models.User, error) {
	//Map Results
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.CreatedUTC, &user.UserName, &user.Email, &user.ProfilePicture); err != nil {
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
	err := row.Scan(&user.ID, &user.CreatedUTC, &user.UserName, &user.Email, &user.ProfilePicture)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("item not found: %w", err)
	}

	return &user, nil
}
