// internal/cache/cache_test.go
package cache

import (
	"testing"
	"time"

	"github.com/HlfDev/pokeapi-wrapper/internal/logger"
	"github.com/HlfDev/pokeapi-wrapper/internal/models"
)

func TestCacheSetAndGet(t *testing.T) {
	logger.Init()
	c := NewCache(1 * time.Second)
	p := models.Pokemon{ID: 1, Name: "Test"}
	c.Set(1, p)
	got, found := c.Get(1)
	if !found {
		t.Fatal("expected found true")
	}
	if got.ID != p.ID {
		t.Fatalf("expected %d got %d", p.ID, got.ID)
	}
	time.Sleep(2 * time.Second)
	_, found = c.Get(1)
	if found {
		t.Fatal("expected found false after expiration")
	}
}
