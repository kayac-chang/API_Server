package cache

import (
	"github.com/patrickmn/go-cache"
)

var storage = cache.New(
	cache.DefaultExpiration,
	cache.DefaultExpiration,
)

type Cache struct {
	*cache.Cache
}

func Get() *Cache {

	return &Cache{storage}
}
