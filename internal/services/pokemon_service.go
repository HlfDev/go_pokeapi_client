package services

import (
	"math/rand"
	"pokeapi/internal/dtos"
	"pokeapi/internal/models"
	"pokeapi/internal/repositories"
)

func GetRandomPokemon() (*models.Pokemon, error) {

	maxPokemonId := 807

	randomID := rand.Intn(maxPokemonId) + 1

	pokemonDTO, err := repositories.FetchPokemonByID(randomID)
	if err != nil {
		return nil, err
	}

	pokemonModel := convertToPokemonModel(*pokemonDTO)
	return &pokemonModel, nil
}

func convertToPokemonModel(dto dtos.PokemonDTO) models.Pokemon {
	var types []string
	for _, t := range dto.Types {
		types = append(types, t.Type.Name)
	}

	var abilities []string
	for _, a := range dto.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}

	return models.Pokemon{
		ID:        dto.ID,
		Name:      dto.Name,
		Height:    dto.Height,
		Weight:    dto.Weight,
		Types:     types,
		Abilities: abilities,
		Sprites: models.Sprites{
			FrontDefault: dto.Sprites.FrontDefault,
			BackDefault:  dto.Sprites.BackDefault,
		},
	}
}
