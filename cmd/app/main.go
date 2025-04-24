package main

import (
	"net/http"

	"github.com/CallumLewisGH/Boondock-Service-Base/database"
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
	//Database Initialisation
	database.GetDB()

	//Creates new server instance
	srv := api.NewServer()

	//Route Registry
	routes.RegisterUserRoutes(srv)
	routes.RegisterShoppingItemRoutes(srv)

	//Starts the server on port 8080
	http.ListenAndServe(":8080", srv)

}

//TODO:
// Complete User Endpoints

// GETBYID
// PUTBYID
// CREATE

// USING THIS CREATE AUTHENTICATION MIDDLEWARE
