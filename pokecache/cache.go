package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	Val       []byte
	createdAt time.Time
}

func (c *Cache) NewCache(interval time.Duration) {
	c.mu.Lock()

	c.data = make(map[string]CacheEntry)
	c.interval = interval

	c.mu.Unlock()

	go func() {
		clk := time.Tick(5 * time.Second)
		for {
			<-clk
			c.mu.Lock()
			for key, val := range c.data {
				if time.Now().Sub(val.createdAt) > c.interval {
					delete(c.data, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}

func (c *Cache) Add(key string, val CacheEntry) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.data == nil {
		return
	}
	val.createdAt = time.Now()
	c.data[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val.Val, ok
}
