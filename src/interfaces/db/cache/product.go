package cache

import (
	"context"
	"github.com/itsmeadi/cart/src/entities/models"
	"sync"
	"time"
)

type ProductCache struct {
	Timeout   time.Duration
	Entry     map[int64]Entry
	ResetTime time.Duration //reset whole cache after this time
	Lock      sync.RWMutex
}

type Entry struct {
	Product models.Product
	SavedAt time.Time
}

func InitProductCache(timeout int64, resetTime int64) *ProductCache {

	prodCache := &ProductCache{
		Timeout:   time.Duration(timeout) * time.Second,
		Entry:     make(map[int64]Entry),
		ResetTime: time.Duration(resetTime) * time.Second,
	}
	go prodCache.ClearCache()
	return prodCache
}

func (cache *ProductCache) SaveCache(ctx context.Context, prod models.Product) {

	cache.Lock.Lock()
	defer cache.Lock.Unlock()
	cache.Entry[prod.ID] = Entry{
		Product: prod,
		SavedAt: time.Now(),
	}
}

func (cache *ProductCache) cacheExpired(timeToCompare time.Time) bool {
	if time.Now().Sub(timeToCompare) > cache.Timeout {
		return false
	}
	return true
}

func (cache *ProductCache) GetCache(ctx context.Context, id int64) models.Product {

	var prod models.Product
	var entry Entry
	cache.Lock.RLock()
	entry = cache.Entry[id]
	cache.Lock.RUnlock()

	if cache.cacheExpired(entry.SavedAt) {
		return prod
	}
	return entry.Product
}

func (cache *ProductCache) ClearCache() {
	for true {
		cache.Entry = make(map[int64]Entry) //reset whole map to free memory
		time.Sleep(cache.ResetTime)
	}
}
