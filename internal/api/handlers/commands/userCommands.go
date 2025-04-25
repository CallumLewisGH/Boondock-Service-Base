package commands

import (
	"fmt"

	dbCommand "github.com/CallumLewisGH/Boondock-Service-Base/database/commands"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	authenticationHelper "github.com/CallumLewisGH/Boondock-Service-Base/internal/helpers/authentication"
)

func CreateUser(apiUserRequest *models.CreateUserRequest) (*models.User, error) {
	passwordHash, err := authenticationHelper.GenerateHash(*apiUserRequest.Password)

	if err != nil {
		return nil, fmt.Errorf("error in hashing %s", err)
	}

	userDBRequest := models.UserDbRequest{
		UserName:     apiUserRequest.UserName,
		PasswordHash: &passwordHash,
		Email:        apiUserRequest.Email,
	}

	user, err := dbCommand.CreateUser(userDBRequest)

	if err != nil {
		return nil, fmt.Errorf("error in execution: %s", err)
	}

	return user, nil
}

func UpdateUserById(model models.UpdateUserRequest) (*models.User, error) {
	userDBRequest := models.UserDbRequest{
		UserName:       model.UserName,
		PasswordHash:   model.PasswordHash,
		Email:          model.Email,
		ProfilePicture: model.ProfilePicture,
		DeletedUTC:     model.DeletedUTC,
		ID:             model.ID,
	}

	user, err := dbCommand.UpdateUserById(userDBRequest)

	if err != nil {
		return nil, fmt.Errorf("error in execution: %s", err)
	}

	return user, nil
}
