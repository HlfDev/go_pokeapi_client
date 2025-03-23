package models

// Pokemon represents a Pokemon entity
type Pokemon struct {
	ID        int      `json:"id" example:"25"`                  // ID of the Pokemon
	Name      string   `json:"name" example:"pikachu"`           // Name
	Height    int      `json:"height" example:"4"`               // Height
	Weight    int      `json:"weight" example:"60"`              // Weight
	Types     []string `json:"types" example:"[\"electric\"]"`   // Types
	Abilities []string `json:"abilities" example:"[\"static\"]"` // Abilities
	Sprites   Sprites  `json:"sprites"`                          // Sprites
}

// Sprites contains Pokemon images
type Sprites struct {
	FrontDefault string `json:"front_default" example:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png"`     // Front sprite
	BackDefault  string `json:"back_default" example:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/25.png"` // Back sprite
}
