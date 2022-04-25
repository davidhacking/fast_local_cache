package local_cache

import (
	"github.com/allegro/bigcache/v3"
	"time"
	"unsafe"
)

type bigCacheDecorator struct {
	cache *bigcache.BigCache
}

var (
	defaultConfig = bigcache.DefaultConfig(20 * time.Minute)
)

// Bytes2Str bytes转string
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func NewBigCaching(config ...bigcache.Config) Caching {
	c := defaultConfig
	if len(config) > 0 {
		c = config[0]
	}
	cache, err := bigcache.NewBigCache(c)
	if err != nil {
		panic(err)
	}
	res := &bigCacheDecorator{
		cache: cache,
	}
	return res
}

func (b *bigCacheDecorator) Get(key []byte) (value []byte, err error) {
	return b.cache.Get(Bytes2Str(key))
}

func (b *bigCacheDecorator) Put(key, value []byte) error {
	return b.cache.Set(Bytes2Str(key), value)
}

func (b *bigCacheDecorator) Del(key []byte) error {
	return b.cache.Delete(Bytes2Str(key))
}
