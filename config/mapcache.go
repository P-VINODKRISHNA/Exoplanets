package config

import (
	"sync"
	"time"
)

var DefaultCache *Caches

func initMapCache() {
	DefaultCache = NewCache(time.Minute * 1440) // Initialize the default cache with a 30-minute expiration time
}

type Caches struct {
	items map[string]Item
	mutex sync.RWMutex
}
type Item struct {
	Value      interface{}
	ExpiryTime time.Time
}

func NewCache(expiration time.Duration) *Caches {
	cache := &Caches{
		items: make(map[string]Item),
	}
	go cache.clearExpired()
	return cache
}
func (c *Caches) Set(key string, data interface{}, expiry time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = Item{
		Value:      data,
		ExpiryTime: time.Now().Add(expiry),
	}
}
func (c *Caches) Get(key string) (map[string]interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, found := c.items[key]
	if !found || item.ExpiryTime.Before(time.Now()) {
		return nil, false
	}
	value, ok := item.Value.(map[string]interface{})
	if !ok {
		// Handle the case where item.Value is not of the expected type
		return nil, false
	}
	return value, true
}
func (c *Caches) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.items, key)
}
func (c *Caches) clearExpired() {
	for {
		time.Sleep(time.Second) // Check every second
		c.mutex.Lock()
		for key, item := range c.items {
			if item.ExpiryTime.Before(time.Now()) {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}
