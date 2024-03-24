package api

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      []byte
	next      string
	prev      string
}

type Cache struct {
	entries      map[string]cacheEntry
	mutex        sync.Mutex
	expiryPeriod time.Duration
}

func NewCache(reapInterval, expiryPeriod time.Duration) *Cache {
	cache := &Cache{
		entries:      make(map[string]cacheEntry),
		expiryPeriod: expiryPeriod,
	}
	go cache.reapLoop(reapInterval)
	return cache
}

func (c *Cache) Add(key string, val []byte, next string, prev string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      val,
		next:      next,
		prev:      prev,
	}
}

func (c *Cache) Get(key string) ([]byte, string, string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, exists := c.entries[key]
	if !exists {
		return nil, "", "", false
	}
	return entry.data, entry.next, entry.prev, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reapEntries()
	}
}

func (c *Cache) reapEntries() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for key, entry := range c.entries {
		if now.Sub(entry.createdAt) > c.expiryPeriod {
			delete(c.entries, key)
		}
	}
}
