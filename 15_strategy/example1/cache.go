package main

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func newCache(evictionAlgo evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: evictionAlgo,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

//func (c *cache) get(key string) {
//	delete(c.storage, key)
//}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}
