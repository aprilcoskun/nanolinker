package db

import (
	"github.com/aprilcoskun/nanolinker/models"
	"time"
)

const defaultMaxCache = 100000

// TODO: Maybe implement  mutex locking?
type Cache struct {
	// mutex *sync.RWMutex
	items map[string]*models.CachedLink
}

func (cache *Cache) Set(url *models.CachedLink) {
	// check if cache limit is reached
	if defaultMaxCache < len(cache.items) {
		for k := range cache.items {
			delete(cache.items, k)
			break
		}
	}

	cache.items[url.ID] = url
}

func (cache *Cache) Get(key string) (link *models.CachedLink, found bool) {
	link = cache.items[key]
	if link == nil {
		found = false
	} else if link.ExpiredAt.Valid && !time.Now().Before(link.ExpiredAt.Time) {
		// Check if url expired
		delete(cache.items, key)

		link = nil
	} else {
		found = true
	}

	return
}

func (cache *Cache) Delete(key string) {
	delete(cache.items, key)
}

var linkCache *Cache
