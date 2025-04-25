package routes

import (
	"net/http"
	"reflect"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/Boondock-Service-Base/internal/api/handlers/queries"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(s *api.Server) {
	s.Router.HandleFunc("/users", getUsers()).Methods(http.MethodGet)
	s.Router.HandleFunc("/users/{id}", getUserById()).Methods(http.MethodGet)
	s.Router.HandleFunc("/users/create/", CreateUser()).Methods(http.MethodPost)
	s.Router.HandleFunc("/users/update/", UpdateUserById()).Methods(http.MethodPut)
}

// GetUsers godoc
// @Summary List all users
// @Description Get all users in the database
// @Tags users
// @Produce json
// @Success 200 {array} models.User "Returns empty array if no users found"
// @Router /users [get]
func getUsers() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		results, err := query.GetAllUsers()

		api.EncodeJsonResponse(w, results, err, http.StatusOK)
	}
	return handler
}

// GetUserById godoc
// @Summary Get a user by id
// @Description Gets a user by UUID
// @Tags users
// @Produce json
// @Param id path string true "User ID" Format(uuid) Example(550e8400-e29b-41d4-a716-446655440000)
// @Success 200 {object} models.User "User found"
// @Failure 400 {string} string "Invalid UUID format"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func getUserById() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		//Decode
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil {
			api.EncodeJsonResponseBadRequest(w, err)
			return
		}

		user, err := query.GetUserById(id)

		if reflect.ValueOf(user).IsNil() {
			api.EncodeJsonResponse(w, nil, err, http.StatusNotFound)
			return
		}

		api.EncodeJsonResponse(w, user, err, http.StatusOK)
	}
	return handler
}

// CreateUser godoc
// @Summary Create a new User
// @Description Creates a new User from provided json data
// @Tags users
// @Produce json
// @Param name body models.CreateUserRequest true "User" Format(models.CreateUserRequest) Example(models.CreateUserRequest)
// @Success 200 {} models.User
// @Router /users/create [post]
func CreateUser() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		model := api.DecodeJsonRequest(r, w, models.CreateUserRequest{})

		user, err := commands.CreateUser(model)

		if reflect.ValueOf(user).IsNil() {
			api.EncodeJsonResponse(w, nil, err, http.StatusConflict)
			return
		}

		api.EncodeJsonResponse(w, user, err, http.StatusOK)

	}
	return handler
}

// UpdateUser godoc
// @Summary Updates a User by id
// @Description Updates a new User with the given id
// @Tags users
// @Produce json
// @Param name body models.UpdateUserRequest true "User" Format(models.UpdateUserRequest) Example(models.UpdateUserRequest)
// @Success 200 {} models.User
// @Router /users/update/ [put]
func UpdateUserById() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {

		model := api.DecodeJsonRequest(r, w, models.UpdateUserRequest{})

		user, err := commands.UpdateUserById(*model)

		if reflect.ValueOf(user).IsNil() {
			api.EncodeJsonResponse(w, nil, err, http.StatusConflict)
			return
		}

		api.EncodeJsonResponse(w, user, err, http.StatusOK)

	}
	return handler
}
