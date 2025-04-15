package api

import (
	"net/http"

	_ "github.com/CallumLewisGH/Boondock-Service-Base/docs" // docs is generated by Swag CLI
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Item represents a shopping item
// @Description Shopping item information
type Item struct {
	Id   uuid.UUID `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Name string    `json:"name" example:"Milk"`
}

// Server holds the router and shopping items
type Server struct {
	*mux.Router
	ShoppingItems []Item
}

// @title Boondock-Service-Base API
// @version 1.0
// @description This is a sample shopping list API
// @host localhost:8080
// @BasePath /
func NewServer() *Server {
	s := &Server{
		Router:        mux.NewRouter(),
		ShoppingItems: []Item{},
	}

	// Swagger
	s.Router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"), //The url pointing to API definition
	))

	return s
}

func (s *Server) HandleFunc(pattern string, handler http.HandlerFunc) {
	s.Router.HandleFunc(pattern, handler)
}
