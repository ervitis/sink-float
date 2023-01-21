package repository

type memcache struct {
}

type Memcache interface {
}

func NewMemcache() Memcache {
	return &memcache{}
}
