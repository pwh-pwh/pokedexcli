package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	internal time.Duration
	cache    map[string]cacheEntry
	mutex    sync.RWMutex
}

func NewCache(internal time.Duration) *PokeCache {
	pokeCache := PokeCache{
		internal: internal,
		cache:    make(map[string]cacheEntry),
		mutex:    sync.RWMutex{},
	}
	go pokeCache.ReapLoop()
	return &pokeCache
}

func (p *PokeCache) Add(key string, val []byte) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.cache[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}
}

func (p *PokeCache) Get(key string) ([]byte, bool) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	entry, ok := p.cache[key]
	return entry.val, ok
}

func (p *PokeCache) ReapLoop() {
	ticker := time.NewTicker(p.internal)
	delKey := make([]string, 0)
	for range ticker.C {
		p.mutex.RLock()
		now := time.Now()
		for key, entry := range p.cache {
			if now.Sub(entry.createAt) > p.internal {
				delKey = append(delKey, key)
			}
		}
		p.mutex.RUnlock()
		p.mutex.Lock()
		for _, dK := range delKey {
			delete(p.cache, dK)
		}
		p.mutex.Unlock()
	}
}

type cacheEntry struct {
	createAt time.Time
	val      []byte
}
