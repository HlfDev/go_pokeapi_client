package dtos

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
