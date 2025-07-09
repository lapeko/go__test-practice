package cache

import "sync"

type Cache[T any] interface {
	Put(string, T)
	Get(string) (T, bool)
	Delete(string)
	Clear()
}

type cache[T any] struct {
	sync.RWMutex
	memory map[string]T
}

func New[T any]() Cache[T] {
	return &cache[T]{memory: make(map[string]T)}
}

func (c *cache[T]) Put(key string, data T) {
	c.Lock()
	defer c.Unlock()
	c.memory[key] = data
}

func (c *cache[T]) Get(key string) (T, bool) {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.memory[key]
	return val, ok
}

func (c *cache[T]) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.memory, key)
}

func (c *cache[T]) Clear() {
	c.Lock()
	defer c.Unlock()
	clear(c.memory)
}
