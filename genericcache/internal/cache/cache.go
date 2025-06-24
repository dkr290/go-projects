package cache

// Cache is the key value store
type Cache[K comparable, V any] struct {
	data map[K]V
}

func New[K comparable, V any]() Cache[K, V] {
	return Cache[K, V]{
		data: make(map[K]V),
	}
}

func (c *Cache[K, V]) Read(key K) (V, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache[K, V]) Upsert(key K, value V) error {
	c.data[key] = value
	return nil
}

// Delete removes the entry from for the givven key
func (c *Cache[K, V]) Delete(key K) {
	delete(c.data, key)
}
