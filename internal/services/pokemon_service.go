package services

import (
	"math/rand"

	"github.com/HlfDev/go_pokeapi_client/internal/models"
	"github.com/HlfDev/go_pokeapi_client/internal/repositories"
)

type PokemonServiceInterface interface {
	GetRandomPokemon() (*models.Pokemon, error)
}

type pokemonService struct {
	repo repositories.PokemonRepositoryInterface
}

func NewPokemonService(repo repositories.PokemonRepositoryInterface) PokemonServiceInterface {
	return &pokemonService{repo: repo}
}

func (s *pokemonService) GetRandomPokemon() (*models.Pokemon, error) {
	maxPokemonID := 1025
	randomID := rand.Intn(maxPokemonID)
	dto, err := s.repo.FetchPokemonByID(randomID)
	if err != nil {
		return nil, err
	}
	pokemon := dto.ToModel()
	return &pokemon, nil
}
