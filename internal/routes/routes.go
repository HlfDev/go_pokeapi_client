package routes

import (
	"github.com/HlfDev/pokeapi-wrapper/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine, handler handlers.PokemonHandler) {
	api := e.Group("/api/v1")
	{
		api.GET("/pokemon/random", handler.GetRandomPokemon)
	}
}
