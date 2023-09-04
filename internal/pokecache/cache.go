package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	ce       map[string]cacheEntry
	mx       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		interval: interval,
		mx:       &sync.Mutex{},
		ce:       map[string]cacheEntry{}}
	go c.reapLoop()
	return c
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mx.Lock()
	cEntry := cacheEntry{createdAt: time.Now(), val: val}
	cache.ce[key] = cEntry
	cache.mx.Unlock()

}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mx.Lock()
	defer cache.mx.Unlock()
	cEntry, ok := cache.ce[key]
	if ok {
		return cEntry.val, true
	}
	return []byte{}, false
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()
	for {
		select {
		case _ = <-ticker.C:
			cache.mx.Lock()
			for key, cEntry := range cache.ce {
				if time.Now().Sub(cEntry.createdAt) > cache.interval {
					cache.mx.Unlock()
					delete(cache.ce, key)
					cache.mx.Lock()
				}
			}
			cache.mx.Unlock()
		}
	}
}
