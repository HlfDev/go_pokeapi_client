package handlers

import (
	"net/http"

	"github.com/HlfDev/pokeapi-wrapper/internal/services"

	"github.com/gin-gonic/gin"
)

type PokemonHandler interface {
	GetRandomPokemon(c *gin.Context)
}

type pokemonHandler struct {
	service services.PokemonServiceInterface
}

func NewPokemonHandler(service services.PokemonServiceInterface) PokemonHandler {
	return &pokemonHandler{service: service}
}

// GetRandomPokemon godoc
// @Summary      Get a random Pokemon
// @Description  Fetches a random Pokemon
// @Tags         pokemon
// @Produce      json
// @Success      200  {object}  models.Pokemon
// @Failure      500  {object}  map[string]string
// @Router       /pokemon/random [get]
func (h *pokemonHandler) GetRandomPokemon(c *gin.Context) {
	pokemon, err := h.service.GetRandomPokemon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pokemon)
}
