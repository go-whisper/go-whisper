package instance

import (
	"sync"

	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable
var onceCache sync.Once

func Cache() *cache2go.CacheTable {
	if cache == nil {
		onceCache.Do(func() {
			cache = cache2go.Cache("common")
		})
	}
	return cache
}
