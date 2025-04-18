package routes

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	models "github.com/CallumLewisGH/Boondock-Service-Base/internal/domain/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func RegisterShoppingItemRoutes(s *api.Server) {
	s.Router.HandleFunc("/items", getShoppingItems(s)).Methods(http.MethodGet)
	s.Router.HandleFunc("/items/", createShoppingItem(s)).Methods(http.MethodPost)
	s.Router.HandleFunc("/items/{id}", getShoppingItemById(s)).Methods(http.MethodGet)
	s.Router.HandleFunc("/items/{id}", deleteShoppingItemById(s)).Methods(http.MethodDelete)
}

// GetItems godoc
// @Summary List all shopping items
// @Description Get full shopping list
// @Tags items
// @Produce json
// @Success 200 {array} models.Item "Returns the shopping list, empty array if no items"
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
// @Success 200 {object} models.Item "Item found"
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
		var item models.Item
		for i := range s.ShoppingItems {
			if s.ShoppingItems[i].Id == id {
				item = s.ShoppingItems[i]
				break
			}
		}
		if models.IsEmpty[models.Item](item) {
			json.NewEncoder(w)
			http.Error(w, reflect.TypeOf(item).Name()+" with id "+idStr+" cannot be found", http.StatusNotFound)

		} else {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

// DeleteItemById godoc
// @Summary Deltete a shopping item by id
// @Description Removes a shopping item from storage by UUID
// @Tags items
// @Produce json
// @Param id path string true "Item ID" Format(uuid) Example(550e8400-e29b-41d4-a716-446655440000)
// @Success 200 {object} models.Item "Item found"
// @Failure 400 {string} string "Invalid UUID format"
// @Failure 404 {string} string "Item not found"
// @Router /items/{id} [delete]
func deleteShoppingItemById(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//Logic
		var item models.Item
		for i := range s.ShoppingItems {
			if s.ShoppingItems[i].Id == id {
				item = s.ShoppingItems[i]
				s.ShoppingItems[i] = s.ShoppingItems[len(s.ShoppingItems)-1]
				s.ShoppingItems = s.ShoppingItems[:len(s.ShoppingItems)-1]
				break
			}
		}
		if models.IsEmpty[models.Item](item) {
			json.NewEncoder(w)
			http.Error(w, reflect.TypeOf(item).Name()+" with id "+idStr+" cannot be found", http.StatusNotFound)

		} else {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

// CreateItem godoc
// @Summary Create a new shopping item
// @Description Creates a new shopping item from provided json data
// @Tags items
// @Produce json
// @Param name body models.Item true "Item" Format(models.Item) Example(models.Item)
// @Success 200 {} models.Item
// @Router /items/ [post]
func createShoppingItem(s *api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Decode
		var i models.Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//Logic
		s.ShoppingItems = append(s.ShoppingItems, i)

		//Encode
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
