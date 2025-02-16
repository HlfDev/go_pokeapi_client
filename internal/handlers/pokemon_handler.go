package handlers

import (
	"net/http"
	"pokeapi/internal/services"

	"github.com/gin-gonic/gin"
)

func GetRandomPokemon(c *gin.Context) {
	pokemon, err := services.GetRandomPokemon()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pokemon)
}
