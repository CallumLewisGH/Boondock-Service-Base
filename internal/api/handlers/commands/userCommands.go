package command

import (
	"fmt"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
)

// GetUsers retrieves all users from the database
func CreateUser(userRequest models.UserRequest) (*models.User, error) {
	// user := models.User{
	// 	UserName = userRequest.UserName
	// 	Email = userRequest.Email
	// }
	user, err := dbCommand.CreateUser()

	if err != nil {
		return nil, fmt.Errorf("failure in execution: %w", err)
	}

	return user, nil

}
