package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"pokeapi/internal/dtos"
)

var pokeAPIBaseURL string

func init() {
	pokeAPIBaseURL = os.Getenv("POKE_API_BASE_URL")

	if pokeAPIBaseURL == "" {
		pokeAPIBaseURL = "https://pokeapi.co/api/v2/pokemon/"
	}

}

func FetchPokemonByID(id int) (*dtos.PokemonDTO, error) {
	url := fmt.Sprintf("%s%d", pokeAPIBaseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha ao buscar Pokémon: status %d", resp.StatusCode)
	}

	var pokemonDTO dtos.PokemonDTO
	if err := json.NewDecoder(resp.Body).Decode(&pokemonDTO); err != nil {
		return nil, err
	}

	return &pokemonDTO, nil
}
