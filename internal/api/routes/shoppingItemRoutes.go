package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	"github.com/google/uuid"
)

func RegisterShoppingItemRoutes(s *api.Server) {
	s.Router.HandleFunc("", getShoppingItems(s)).Methods(http.MethodGet)
	s.Router.HandleFunc("/", createShoppingItem(s)).Methods(http.MethodPost)
	s.Router.HandleFunc("/{id}", getShoppingItemById(s)).Methods(http.MethodGet)
}

// GetItems godoc
// @Summary List all shopping items
// @Description Get full shopping list
// @Tags items
// @Produce json
// @Success 200 {array} api.Item
// @Router /items [get]
func getShoppingItems(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Encode
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.ShoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getShoppingItemById(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		var id uuid.UUID
		if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		//Logic
		var item api.Item
		for i := range s.ShoppingItems {
			if s.ShoppingItems[i].Id == id {
				item = s.ShoppingItems[i]
			}
		}
		//Encode
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func createShoppingItem(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		var i api.Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Logic
		i.Id = uuid.New()
		s.ShoppingItems = append(s.ShoppingItems, i)

		//Encode
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
