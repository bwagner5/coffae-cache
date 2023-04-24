package simplecounter

import "sync"

type SimpleCounter[K comparable] struct {
	counts map[K]uint
	mu     sync.RWMutex
}

func NewSimpleCounter[K comparable]() *SimpleCounter[K] {
	return &SimpleCounter[K]{
		counts: map[K]uint{},
	}
}

func (c *SimpleCounter[K]) Count(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key] += 1
}

func (c *SimpleCounter[K]) Get(key K) uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counts[key]
}
