package shop

type cacher interface {
	Clear()
	Get(string) (int, bool)
	Set(hash string, sum int)
}

type bigCache map[string]int

func NewBigCache(cap int) *bigCache {
	c := bigCache(make(map[string]int, cap))
	return &c
}

func (c *bigCache) Clear() {
	*c = make(map[string]int, 100)
}

func (c bigCache) Get(hash string) (int, bool) {
	v, ok := c[hash]
	return v, ok
}

func (c *bigCache) Set(hash string, sum int) {
	(*c)[hash] = sum
}