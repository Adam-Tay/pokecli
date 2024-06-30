package pokecache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	values   map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	Value     []byte
}

func NewCache(interval time.Duration) Cache {
	//ticker := time.NewTicker(interval)

	c := Cache{
		values:   make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	c.mu.Lock()
	defer c.mu.Unlock()
	for {
		<-ticker.C
		for index, item := range c.values {
			if item.createdAt.After(time.Now().Add(c.interval)) {
				delete(c.values, index)
			}
		}
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var entry cacheEntry = cacheEntry{
		createdAt: time.Now(),
		Value:     val,
	}
	c.values[key] = entry
}

func (c *Cache) Get(key string) (cacheEntry, error) {
	ret, ok := c.values[key]
	if ok {
		return ret, nil
	}
	return ret, errors.New("entry does not exist")
}
