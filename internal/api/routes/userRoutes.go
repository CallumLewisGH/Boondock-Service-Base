package routes

import (
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	query "github.com/CallumLewisGH/Boondock-Service-Base/internal/api/handlers/queries"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(s *api.Server) {
	s.Router.HandleFunc("/users", getUsers()).Methods(http.MethodGet)
	s.Router.HandleFunc("/users/{id}", getUserById()).Methods(http.MethodGet)
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

		api.JsonResponse(w, results, err, http.StatusOK)
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
			api.JsonResponseBadRequest(w, err)
			return
		}

		user, err := query.GetUserById(id)

		api.JsonResponse(w, user, err, http.StatusOK)
	}
	return handler
}

// // DeleteUserById godoc
// // @Summary Deltete a shopping User by id
// // @Description Removes a shopping User from storage by UUID
// // @Tags Users
// // @Produce json
// // @Param id path string true "User ID" Format(uuid) Example(550e8400-e29b-41d4-a716-446655440000)
// // @Success 200 {object} models.User "User found"
// // @Failure 400 {string} string "Invalid UUID format"
// // @Failure 404 {string} string "User not found"
// // @Router /Users/{id} [delete]
// func deleteUserById(s *api.Server) http.HandlerFunc {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		//Decode
// 		idStr := mux.Vars(r)["id"]
// 		id, err := uuid.Parse(idStr)

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 		}

// 		//Logic
// 		var User models.User
// 		for i := range s.Users {
// 			if s.Users[i].Id == id {
// 				User = s.Users[i]
// 				s.Users[i] = s.Users[len(s.Users)-1]
// 				s.Users = s.Users[:len(s.Users)-1]
// 				break
// 			}
// 		}
// 		if models.IsEmpty[models.User](User) {
// 			json.NewEncoder(w)
// 			http.Error(w, reflect.TypeOf(User).Name()+" with id "+idStr+" cannot be found", http.StatusNotFound)

// 		} else {
// 			if err := json.NewEncoder(w).Encode(User); err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 	}
// 	return handler
// }

// // CreateUser godoc
// // @Summary Create a new shopping User
// // @Description Creates a new shopping User from provided json data
// // @Tags Users
// // @Produce json
// // @Param name body models.User true "User" Format(models.User) Example(models.User)
// // @Success 200 {} models.User
// // @Router /Users/ [post]
// func createUser(s *api.Server) http.HandlerFunc {
// 	handler := func(w http.ResponseWriter, r *http.Request) {
// 		//Decode
// 		var i models.User
// 		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		//Logic
// 		s.Users = append(s.Users, i)

// 		//Encode
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(i); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	return handler
// }
