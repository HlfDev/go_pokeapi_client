package dtos

import (
	"testing"
)

func TestToModel(t *testing.T) {
	typeVal := struct {
		Name string `json:"name"`
	}{Name: "grass"}
	abilityVal := struct {
		Name string `json:"name"`
	}{Name: "overgrow"}
	dto := &PokemonDTO{
		ID:     1,
		Name:   "bulbasaur",
		Height: 7,
		Weight: 69,
		Types: []TypeDTO{
			{Slot: 1, Type: typeVal},
		},
		Abilities: []AbilityDTO{
			{Ability: abilityVal},
		},
		Sprites: SpritesDTO{
			FrontDefault: "front.png",
			BackDefault:  "back.png",
		},
	}
	m := dto.ToModel()
	if m.ID != dto.ID || m.Name != dto.Name || m.Height != dto.Height || m.Weight != dto.Weight {
		t.Fatalf("conversion error: %+v", m)
	}
	if len(m.Types) != 1 || m.Types[0] != "grass" {
		t.Fatalf("types conversion error: %+v", m.Types)
	}
	if len(m.Abilities) != 1 || m.Abilities[0] != "overgrow" {
		t.Fatalf("abilities conversion error: %+v", m.Abilities)
	}
	if m.Sprites.FrontDefault != "front.png" || m.Sprites.BackDefault != "back.png" {
		t.Fatalf("sprites conversion error: %+v", m.Sprites)
	}
}
