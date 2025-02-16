package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"pokeapi/internal/models"

	"github.com/gin-gonic/gin"
)

type mockService struct {
	pokemon *models.Pokemon
	err     error
}

func (m *mockService) GetRandomPokemon() (*models.Pokemon, error) {
	return m.pokemon, m.err
}

func TestGetRandomPokemon_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mSvc := &mockService{
		pokemon: &models.Pokemon{
			ID:        1,
			Name:      "bulbasaur",
			Height:    7,
			Weight:    69,
			Types:     []string{"grass"},
			Abilities: []string{"overgrow"},
			Sprites:   models.Sprites{FrontDefault: "front.png", BackDefault: "back.png"},
		},
		err: nil,
	}
	handler := NewPokemonHandler(mSvc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/v1/pokemon/random", nil)
	handler.GetRandomPokemon(c)
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}

func TestGetRandomPokemon_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mSvc := &mockService{
		pokemon: nil,
		err:     errors.New("service error"),
	}
	handler := NewPokemonHandler(mSvc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/api/v1/pokemon/random", nil)
	handler.GetRandomPokemon(c)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}
}
