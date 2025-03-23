package cache

import (
	"sync"
	"time"

	"github.com/HlfDev/pokeapi-wrapper/internal/logger"
	"github.com/HlfDev/pokeapi-wrapper/internal/models"

	"go.uber.org/zap"
)

type CacheEntry struct {
	Pokemon    models.Pokemon
	Expiration time.Time
}

type Cache struct {
	mu    sync.RWMutex
	store map[int]CacheEntry
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		store: make(map[int]CacheEntry),
		ttl:   ttl,
	}
}

func (c *Cache) Get(id int) (models.Pokemon, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.store[id]
	if !exists || time.Now().After(entry.Expiration) {
		logger.Log.Info("cache miss", zap.Int("id", id))
		return models.Pokemon{}, false
	}
	logger.Log.Info("cache hit", zap.Int("id", id))
	return entry.Pokemon, true
}

func (c *Cache) Set(id int, pokemon models.Pokemon) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[id] = CacheEntry{
		Pokemon:    pokemon,
		Expiration: time.Now().Add(c.ttl),
	}
	logger.Log.Info("cache set", zap.Int("id", id), zap.Duration("ttl", c.ttl))
}
