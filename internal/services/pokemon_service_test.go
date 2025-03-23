package services

import (
	"errors"
	"testing"

	"github.com/HlfDev/pokeapi-wrapper/internal/dtos"
)

type mockRepo struct {
	dto *dtos.PokemonDTO
	err error
}

func (m *mockRepo) FetchPokemonByID(id int) (*dtos.PokemonDTO, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.dto, nil
}

func TestGetRandomPokemon(t *testing.T) {
	dto := &dtos.PokemonDTO{
		ID:     1,
		Name:   "Test",
		Height: 10,
		Weight: 100,
		Sprites: dtos.SpritesDTO{
			FrontDefault: "front.png",
			BackDefault:  "back.png",
		},
		Types:     []dtos.TypeDTO{},
		Abilities: []dtos.AbilityDTO{},
	}
	mock := &mockRepo{dto: dto, err: nil}
	service := NewPokemonService(mock)
	pokemon, err := service.GetRandomPokemon()
	if err != nil {
		t.Fatal(err)
	}
	if pokemon.Name != "Test" {
		t.Fatalf("expected Test got %s", pokemon.Name)
	}
}

func TestGetRandomPokemonError(t *testing.T) {
	mock := &mockRepo{dto: nil, err: errors.New("error")}
	service := NewPokemonService(mock)
	_, err := service.GetRandomPokemon()
	if err == nil {
		t.Fatal("expected error")
	}
}
