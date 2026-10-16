package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}
	c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, exists := c.entries[key]
	if exists {
		return item.val, exists
	}
	return []byte{}, false
}

func (c *Cache) GetAll() map[string]cacheEntry {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entries
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.createdAt) > interval {
		}
		delete(c.entries, key)
	}
}
