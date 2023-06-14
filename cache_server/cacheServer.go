package cache_server

import (
	"github.com/VINDA-98/memCache/cache"
	"time"
)

// @Title  cache_server
// @Description  MyGO
// @Author  WeiDa  2023/6/14 12:26
// @Update  WeiDa  2023/6/14 12:26

type CacheServer struct {
	mCache cache.Cache
}

func NewCacheServer() *CacheServer {
	c := &CacheServer{
		cache.NewMemCache(),
	}
	return c
}

func (c *CacheServer) SetMaxMemory(size string) bool {
	c.mCache.SetMaxMemory(size)
	return true
}

func (c *CacheServer) Set(key string, val interface{}, expire ...time.Duration) {
	tmpExpire := 0 * time.Second
	if len(expire) > 0 {
		tmpExpire = expire[0]
	}
	c.mCache.Set(key, val, tmpExpire)
}

func (c *CacheServer) Get(key string) (interface{}, bool) {
	return c.mCache.Get(key)
}

func (c *CacheServer) Del(key string) bool {
	return c.mCache.Del(key)
}

func (c *CacheServer) Exists(key string) bool {
	return c.mCache.Exists(key)
}

func (c *CacheServer) Flush() bool {
	return c.mCache.Flush()
}

func (c *CacheServer) Keys() int64 {
	return c.mCache.Keys()
}
