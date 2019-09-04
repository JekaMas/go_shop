package shop

type BigCacheStub struct{}

func (b *BigCacheStub) Clear() {}

func (b BigCacheStub) Get(string) (int, bool) {
	return 0, false
}

func (b *BigCacheStub) Set(hash string, sum int) {}
