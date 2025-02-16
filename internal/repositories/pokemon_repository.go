package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pokeapi/internal/dtos"
)

type PokemonRepository interface {
	FetchPokemonByID(id int) (*dtos.PokemonDTO, error)
}

type pokemonRepository struct {
	client  *http.Client
	baseURL string
}

func NewPokemonRepository(client *http.Client) PokemonRepository {
	baseURL := os.Getenv("POKE_API_BASE_URL")
	if baseURL == "" {
		baseURL = "https://pokeapi.co/api/v2/pokemon/"
	}
	return &pokemonRepository{client: client, baseURL: baseURL}
}

func (r *pokemonRepository) FetchPokemonByID(id int) (*dtos.PokemonDTO, error) {
	url := fmt.Sprintf("%s%d", r.baseURL, id)
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch Pokémon: status %d", resp.StatusCode)
	}

	var pokemonDTO dtos.PokemonDTO
	if err := json.NewDecoder(resp.Body).Decode(&pokemonDTO); err != nil {
		return nil, err
	}

	return &pokemonDTO, nil
}
