package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	shopping "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models/shopping"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterShoppingItemRoutes(s *api.Server) {
	s.Router.HandleFunc("/items", getShoppingItems(s)).Methods(http.MethodGet)
	s.Router.HandleFunc("/items/", createShoppingItem(s)).Methods(http.MethodPost)
	s.Router.HandleFunc("/items/{id}", getShoppingItemById(s)).Methods(http.MethodGet)
}

// GetItems godoc
// @Summary List all shopping items
// @Description Get full shopping list
// @Tags items
// @Produce json
// @Success 200 {array} shopping.Item "Returns the shopping list, empty array if no items"
// @Router /items [get]
func getShoppingItems(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.ShoppingItems); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// GetItemById godoc
// @Summary Get a shopping item by id
// @Description Gets a shopping item by UUID
// @Tags items
// @Produce json
// @Param id path string true "Item ID" Format(uuid) Example(550e8400-e29b-41d4-a716-446655440000)
// @Success 200 {object} shopping.Item "Item found"
// @Failure 400 {string} string "Invalid UUID format"
// @Failure 404 {string} string "Item not found"
// @Router /items/{id} [get]
func getShoppingItemById(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//Logic
		var item shopping.Item
		for i := range s.ShoppingItems {
			if s.ShoppingItems[i].Id == id {
				item = s.ShoppingItems[i]
			}
		}
		var emptyItem shopping.Item
		if item == emptyItem {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		//Encode
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// CreateItem godoc
// @Summary Create a new shopping item
// @Description Creates a new shopping item from provided json data
// @Tags items
// @Produce json
// @Param name body shopping.Item true "Item" Format(shopping.Item) Example(shopping.Item)
// @Success 200 {} shopping.Item
// @Router /items/ [post]
func createShoppingItem(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		var i shopping.Item
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
