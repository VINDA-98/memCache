package cache

import (
	"log"
	"sync"
	"time"
)

// @Title  cache
// @Description  MyGO
// @Author  WeiDa  2023/6/14 10:50
// @Update  WeiDa  2023/6/14 10:50

type memCacheValue struct {
	value interface{}
	//插入时间
	insertTime time.Time
	//有效时长
	expire time.Duration
	//value 大小
	size int64
}

// MemCache 实现内存管理的结构体，实现Cache接口
type MemCache struct {
	//maxMemorySizeStr 最大内存
	maxMemorySizeStr string
	//maxMemorySize 最大内存对应的int64
	maxMemorySize int64
	//currMemorySize 当前使用内存大小
	currMemorySize int64
	//缓存键值对值
	values map[string]*memCacheValue
	//读写锁
	locker sync.RWMutex
	//清除过期缓存时间间隔
	clearExpiredItemTimeInterval time.Duration
}

// NewMemCache 创建MemCache实例
func NewMemCache() *MemCache {
	_memCache := &MemCache{
		values:                       make(map[string]*memCacheValue, 0),
		clearExpiredItemTimeInterval: time.Second,
	}
	go _memCache.clearExpiredItem()
	return _memCache
}

// SetMaxMemory 设置缓存最大内存
func (m *MemCache) SetMaxMemory(size string) bool {
	m.maxMemorySize, m.maxMemorySizeStr = ParseSize(size)
	return true
}

// Set 设置缓存，expire后过期
func (m *MemCache) Set(key string, val interface{}, expire time.Duration) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	valSize := GetValSize(val)
	mcv := &memCacheValue{
		value:      val,
		insertTime: time.Time{},
		size:       valSize,
		expire:     expire,
	}
	m.del(key)
	m.add(key, mcv)
	if m.currMemorySize > m.maxMemorySize {
		m.del(key)
		log.Fatalf("max memory size %s", m.maxMemorySizeStr)
	}
	return true
}

// Get 获取缓存
func (m *MemCache) Get(key string) (interface{}, bool) {
	m.locker.Lock()
	defer m.locker.Unlock()
	//得到key对应的value
	mcv, b := m.get(key)
	if !b {
		return nil, false
	}
	//判定缓存是否过期，过期则删除该缓存项
	if mcv.expire > 0 && time.Now().Sub(mcv.insertTime) > mcv.expire {
		m.del(key)
		return nil, false
	}
	return mcv.value, true
}

// Del 删除缓存
func (m *MemCache) Del(key string) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.del(key)
	return true
}

// Exists 判断缓存是否存在
func (m *MemCache) Exists(key string) bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	_, b := m.get(key)
	return b
}

// Flush 清空所有缓存
func (m *MemCache) Flush() bool {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.values = make(map[string]*memCacheValue, 0)
	m.currMemorySize = 0
	return true
}

// Keys	获取缓存中所有key的数量
func (m *MemCache) Keys() int64 {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return int64(len(m.values))
}

// clearExpiredItem 清除过期缓存
func (m *MemCache) clearExpiredItem() {
	timeTicker := time.NewTicker(m.clearExpiredItemTimeInterval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, value := range m.values {
				// 判断当前时间是否在指定时间之后
				if value.expire != 0 && time.Now().After(value.insertTime.Add(value.expire)) {
					m.locker.Lock()
					m.del(key)
					m.locker.Unlock()
				}
			}
		}
	}
}

// get 获取key对应的value
func (m *MemCache) get(key string) (*memCacheValue, bool) {
	val, ok := m.values[key]
	return val, ok
}

// del 删除key
func (m *MemCache) del(key string) {
	//得到key对应的value
	value, b := m.get(key)
	//删除key delete:删除map中的元素
	delete(m.values, key)
	if b && value != nil {
		m.currMemorySize -= value.size
	}
}

// add 添加key
func (m *MemCache) add(key string, mcv *memCacheValue) {
	//添加key
	m.values[key] = mcv
	//增加当前内存大小
	m.currMemorySize += mcv.size
}
