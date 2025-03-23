package repositories

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/HlfDev/go_pokeapi_client/internal/cache"
	"github.com/HlfDev/go_pokeapi_client/internal/logger"
)

func init() {
	logger.Init()
}

func TestFetchPokemonByID_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonResponse := `{
			"id": 1,
			"name": "bulbasaur",
			"height": 7,
			"weight": 69,
			"types": [{"slot": 1, "type": {"name": "grass"}}],
			"abilities": [{"ability": {"name": "overgrow"}}],
			"sprites": {"front_default": "front.png", "back_default": "back.png"}
		}`
		w.Write([]byte(jsonResponse))
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := cache.NewCache(10 * time.Minute)
	repo := &pokemonRepository{
		client:  httpClient,
		cache:   c,
		baseURL: ts.URL + "/",
	}
	dto, err := repo.FetchPokemonByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if dto.ID != 1 || dto.Name != "bulbasaur" {
		t.Fatalf("unexpected dto: %+v", dto)
	}
}

func TestFetchPokemonByID_ErrorStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := cache.NewCache(10 * time.Minute)
	repo := &pokemonRepository{
		client:  httpClient,
		cache:   c,
		baseURL: ts.URL + "/",
	}
	_, err := repo.FetchPokemonByID(1)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestFetchPokemonByID_InvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer ts.Close()
	httpClient := ts.Client()
	c := cache.NewCache(10 * time.Minute)
	repo := &pokemonRepository{
		client:  httpClient,
		cache:   c,
		baseURL: ts.URL + "/",
	}
	_, err := repo.FetchPokemonByID(1)
	if err == nil {
		t.Fatal("expected error for invalid json")
	}
}
