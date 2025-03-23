package models

// Pokemon represents a Pokemon entity with its basic information
type Pokemon struct {
	ID        int      `json:"id" example:"25"`                  // The unique identifier for the Pokemon
	Name      string   `json:"name" example:"pikachu"`           // The name of the Pokemon
	Height    int      `json:"height" example:"4"`               // Height in decimeters
	Weight    int      `json:"weight" example:"60"`              // Weight in hectograms
	Types     []string `json:"types" example:"[\"electric\"]"`   // Types of the Pokemon
	Abilities []string `json:"abilities" example:"[\"static\"]"` // Abilities of the Pokemon
	Sprites   Sprites  `json:"sprites"`                          // Visual representations of the Pokemon
}

// Sprites contains URLs to Pokemon images
type Sprites struct {
	FrontDefault string `json:"front_default" example:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png"`     // URL to the front default sprite
	BackDefault  string `json:"back_default" example:"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/25.png"` // URL to the back default sprite
}
