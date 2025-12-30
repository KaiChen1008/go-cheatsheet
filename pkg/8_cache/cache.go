package threadsafe

import "sync"

// use sync.Map

func main() {
	m := sync.Map{}
	m.LoadOrStore("key", "block") // use load or store to handle concurrency issue
}

type Cache struct {
	mu   sync.Mutex
	data map[string]string
}

func NewMap() *Cache {
	return &Cache{
		mu:   sync.Mutex{},
		data: make(map[string]string),
	}
}

func (c *Cache) Get(k string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.data[k]
	return v, ok
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func (c *Cache) LoadOrStore(k, v string) string {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v, ok := c.data[k]; ok {
		return v
	}
	c.data[k] = v
	return v
}
