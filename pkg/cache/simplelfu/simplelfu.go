package simplelfu

import (
	"github.com/bwagner5/coffee-cache/pkg/counter"
	"github.com/bwagner5/coffee-cache/pkg/counter/simplercounter"
)

type SimpleLFUCache[K comparable, V any] struct {
	sorted  []*V
	lookup  map[K]*V
	counter counter.Counter[K]
}

func NewSimpleLFUCache[K comparable, V any]() *SimpleLFUCache[K, V] {
	return &SimpleLFUCache[K, V]{
		sorted: []*V{},
		lookup: map[K]*V{},
		counter: simplercounter.NewSimpleCounter[K](),
	}
}

func (lfu SimpleLFUCache[K, V]) Put(key K, val V) {
	lfu.counter.
}

func (lfu SimpleLFUCache[K, V]) Get(key K) V {
	return
}
