package bootstrap

import (
	"log"
	"net/http"
	"pokeapi/internal/handlers"
	"pokeapi/internal/repositories"
	"pokeapi/internal/routes"
	"pokeapi/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	httpClient := &http.Client{}
	repo := repositories.NewPokemonRepository(httpClient)
	service := services.NewPokemonService(repo)
	handler := handlers.NewPokemonHandler(service)

	e := gin.Default()

	routes.SetupRoutes(e, handler)

	port := ":8080"
	log.Printf("Starting server on %s...", port)
	if err := e.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
