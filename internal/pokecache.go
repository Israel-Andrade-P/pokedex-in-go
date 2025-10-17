package pokecache

import (
	"sync"
	"time"
)

type (
	Cache struct {
		pokeCache map[string]CacheEntry
		mu        sync.Mutex
	}
	CacheEntry struct {
		createdAt time.Time
		val       []byte
	}
)

func NewCache(interval time.Duration) Cache {

	return Cache{
		pokeCache: make(map[string]CacheEntry),
		mu:        sync.Mutex{},
	}
}

func (cache *Cache) Add(key string, data []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.pokeCache[key] = CacheEntry{createdAt: time.Now(), val: data}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	for k, entry := range cache.pokeCache {
		if k == key {
			return entry.val, true
		}
	}
	return nil, false
}
