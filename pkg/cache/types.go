package cache

// Cache is the public interface for a cache
type Cache[K comparable, V any] interface {
	Get(K) V
	Put(K, V)
}
