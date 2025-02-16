package dtos

import "pokeapi/internal/models"

type PokemonDTO struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Height    int          `json:"height"`
	Weight    int          `json:"weight"`
	Types     []TypeDTO    `json:"types"`
	Abilities []AbilityDTO `json:"abilities"`
	Sprites   SpritesDTO   `json:"sprites"`
}

type TypeDTO struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type AbilityDTO struct {
	Ability struct {
		Name string `json:"name"`
	} `json:"ability"`
}

type SpritesDTO struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}

func (dto *PokemonDTO) ToModel() models.Pokemon {
	types := make([]string, len(dto.Types))
	for i, t := range dto.Types {
		types[i] = t.Type.Name
	}

	abilities := make([]string, len(dto.Abilities))
	for i, a := range dto.Abilities {
		abilities[i] = a.Ability.Name
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
