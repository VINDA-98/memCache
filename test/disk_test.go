package test

import (
	"github.com/VINDA-98/memCache/cache"
	"math"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

// @Title  test
// @Description  MyGO
// @Author  WeiDa  2023/6/14 17:19
// @Update  WeiDa  2023/6/14 17:19

func TestDisk(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"vinda1", 678, time.Second * 10},
		{"vinda2", false, time.Second * 11},
		{"vinda3", true, time.Second * 12},
		{"data", map[string]interface{}{"a": 3}, time.Second * 13},
		{"vinda111", "vinda111", time.Second * 14},
		{"vinda222", "这里是字符串", time.Second * 15},
		{"vinda333", "这里是字符串这里是字符串这里是字符串这里是字符串这里是字", time.Second * 15},
	}
	c := cache.NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testData {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		if item.key != "data" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "data" && !ok1 {
			t.Error("缓存取值数据与预期不一致")
		}
	}
	if int64(len(testData)) != c.Keys() {
		t.Error("缓存数量不一致")
	}
	c.Del(testData[0].key)
	c.Del(testData[1].key)

	if int64(len(testData)) != c.Keys()+2 {
		t.Error("缓存数量不一致")
	}
	time.Sleep(time.Second * 16)
	if c.Keys() != 0 {
		t.Error("过期缓存数据清空失败")
	}
}

func TestDiskLimit(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("最大内存限制测试不通过")
		}
	}()
	testStr := "这里是字符串这里是字符串这里是字符串这里是字符串这里是字"
	c := cache.NewMemCache()
	c.SetMaxMemory("1KB")

	valSize := int64(unsafe.Sizeof(testStr))
	num := int(math.Ceil(1024 / float64(valSize)))

	for i := 0; i <= num; i++ {
		c.Set(strconv.Itoa(i), testStr, time.Second*30)
	}
}
