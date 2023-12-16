package db

import (
	lru "github.com/hashicorp/golang-lru"
)

const (
	CacheLimit = 10000
)

type Cache struct {
	logCache *lru.Cache
}

func NewCache() *Cache {
	logCache, _ := lru.New(CacheLimit)

	return &Cache{
		logCache: logCache,
	}
}
