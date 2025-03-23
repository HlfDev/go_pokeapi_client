package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HlfDev/go_pokeapi_client/internal/cache"
	"github.com/HlfDev/go_pokeapi_client/internal/dtos"
	"github.com/HlfDev/go_pokeapi_client/internal/logger"

	"go.uber.org/zap"
)

type PokemonRepositoryInterface interface {
	FetchPokemonByID(id int) (*dtos.PokemonDTO, error)
}

type pokemonRepository struct {
	client  *http.Client
	cache   *cache.Cache
	baseURL string
}

func NewPokemonRepository(client *http.Client, cache *cache.Cache) PokemonRepositoryInterface {
	baseURL := "https://pokeapi.co/api/v2/pokemon/"
	return &pokemonRepository{client: client, cache: cache, baseURL: baseURL}
}

func (r *pokemonRepository) FetchPokemonByID(id int) (*dtos.PokemonDTO, error) {
	if pokemon, found := r.cache.Get(id); found {
		dto := dtos.PokemonDTO{
			ID:     pokemon.ID,
			Name:   pokemon.Name,
			Height: pokemon.Height,
			Weight: pokemon.Weight,
			Sprites: dtos.SpritesDTO{
				FrontDefault: pokemon.Sprites.FrontDefault,
				BackDefault:  pokemon.Sprites.BackDefault,
			},
		}
		logger.Log.Info("repository cache hit", zap.Int("id", id))
		return &dto, nil
	}
	url := fmt.Sprintf("%s%d", r.baseURL, id)
	logger.Log.Info("repository API call", zap.Int("id", id), zap.String("url", url))
	resp, err := r.client.Get(url)
	if err != nil {
		logger.Log.Error("repository API error", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status %d", resp.StatusCode)
		logger.Log.Error("repository API status error", zap.Error(err))
		return nil, err
	}
	var dto dtos.PokemonDTO
	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		logger.Log.Error("repository decode error", zap.Error(err))
		return nil, err
	}
	pokemonModel := dto.ToModel()
	r.cache.Set(id, pokemonModel)
	logger.Log.Info("repository set cache", zap.Int("id", id))
	return &dto, nil
}
