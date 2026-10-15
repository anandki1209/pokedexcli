package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	cache := Cache{
		interval: interval,
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

	go cache.reapLoop()
	return &cache

}

// use pointer reference if you want to have shared data structure
func (c *Cache) Add(key string, val []byte) {
	chEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = chEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	ch, ok := c.cacheMap[key]
	if !ok {
		return []byte{}, false
	}
	return ch.val, true
}

// lock the whole map while reaping
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(time.Now().UTC(), c.interval)

	}
}

// reap takes time and interval an d delete all expired entries in map
func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, ch := range c.cacheMap {

		if ch.createdAt.Before(now.Add(-interval)) {
			delete(c.cacheMap, k)
		}
	}
}
