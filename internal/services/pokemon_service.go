package services

import (
	"math/rand"
	"pokeapi/internal/models"
	"pokeapi/internal/repositories"
)

type PokemonService interface {
	GetRandomPokemon() (*models.Pokemon, error)
}

type pokemonService struct {
	repo repositories.PokemonRepository
}

func NewPokemonService(repo repositories.PokemonRepository) PokemonService {
	return &pokemonService{repo: repo}
}

func (s *pokemonService) GetRandomPokemon() (*models.Pokemon, error) {
	maxPokemonId := 1025
	randomID := rand.Intn(maxPokemonId)

	pokemonDTO, err := s.repo.FetchPokemonByID(randomID)
	if err != nil {
		return nil, err
	}

	pokemonModel := pokemonDTO.ToModel()
	return &pokemonModel, nil
}
