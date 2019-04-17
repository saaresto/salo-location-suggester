package cache

import (
	"sync"
)

type Cache struct {
	mutex  sync.RWMutex
	values map[string]interface{}
}

func InitializeCache() *Cache {
	return &Cache{
		values: make(map[string]interface{}),
	}
}

func (c *Cache) PutValue(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.values[key] = value
}

func (c *Cache) GetValue(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok := c.values[key]
	return value, ok
}

func (c *Cache) DeleteValue(key string) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	delete(c.values, key)
}

func (c Cache) GetSize() int {
	return len(c.values)
}
