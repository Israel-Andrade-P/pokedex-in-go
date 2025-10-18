package pokecache

import (
	"sync"
	"time"
)

type (
	Cache struct {
		pokeCache map[string]CacheEntry
		mu        sync.Mutex
		interval  time.Duration
	}
	CacheEntry struct {
		createdAt time.Time
		val       []byte
	}
)

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		pokeCache: make(map[string]CacheEntry),
		mu:        sync.Mutex{},
		interval:  interval,
	}

	go c.reapLoop()

	return c
}

func (cache *Cache) Add(key string, data []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.pokeCache[key] = CacheEntry{createdAt: time.Now(), val: data}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, exists := cache.pokeCache[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		cache.reap()
	}
}

func (cache *Cache) reap() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	now := time.Now()
	for key, entry := range cache.pokeCache {
		if now.Sub(entry.createdAt) > cache.interval {
			delete(cache.pokeCache, key)
		}
	}
}
