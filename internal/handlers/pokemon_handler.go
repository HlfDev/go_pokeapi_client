package handlers

import (
	"net/http"
	"pokeapi/internal/services"

	"github.com/gin-gonic/gin"
)

type PokemonHandler interface {
	GetRandomPokemon(c *gin.Context)
}

type pokemonHandler struct {
	service services.PokemonService
}

func NewPokemonHandler(service services.PokemonService) PokemonHandler {
	return &pokemonHandler{service: service}
}

func (h *pokemonHandler) GetRandomPokemon(c *gin.Context) {
	pokemon, err := h.service.GetRandomPokemon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pokemon)
}
