package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	query "github.com/CallumLewisGH/Boondock-Service-Base/internal/api/handlers/queries"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api/middleware"
	repository "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/repositories"
)

func RegisterUserRoutes(s *api.Server) {
	s.Router.HandleFunc("/users", getUsers()).Methods(http.MethodGet)
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

		w.Header().Set("Content-Type", "application/json")

		results, err := repository.Query(query.GetUsers)

		if err != nil {
			json.NewEncoder(w)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(results); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return middleware.HTTPLogger(handler)
}
