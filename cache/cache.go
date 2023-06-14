package cache

import "time"

// @Title  cache
// @Description  MyGO
// @Author  WeiDa  2023/6/14 10:44
// @Update  WeiDa  2023/6/14 10:44

type Cache interface {
	// SetMaxMemory size : 1KB 100KB 1MB 2MB 1GB
	SetMaxMemory(size string) bool
	// Set 将value写入缓存
	Set(key string, val interface{}, expire time.Duration) bool
	// Get 根据key值获取value
	Get(key string) (interface{}, bool)
	// Del 删除key值
	Del(key string) bool
	// Exists 判断key是否存在
	Exists(key string) bool
	// Flush 清空所有key
	Flush() bool
	// Keys 获取缓存中所有key的数量
	Keys() int64
}
