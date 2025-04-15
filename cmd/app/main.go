package main

import (
	"fmt"
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api"
	"github.com/CallumLewisGH/Boondock-Service-Base/internal/api/routes"
	_ "github.com/Masterminds/squirrel"
	_ "github.com/joho/godotenv"
	_ "github.com/lann/builder"
)

// @title Boondock-Service-Base API
// @version 1.0
// @description This is the base service for boondocks API
// @host localhost:8080
// @BasePath /
func main() {
	//Entry Point
	fmt.Println("Server Starting!")
	srv := api.NewServer()

	routes.RegisterShoppingItemRoutes(srv)

	http.ListenAndServe(":8080", srv)

}
