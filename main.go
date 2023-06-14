package main

import (
	"github.com/VINDA-98/memCache/cache_server"
	"log"
)

// @Title  memCache
// @Description  MyGO
// @Author  WeiDa  2023/6/14 12:24
// @Update  WeiDa  2023/6/14 12:24

func main() {
	cache := cache_server.NewCacheServer()
	cache.SetMaxMemory("1MB")
	cache.Set("name", "weida")
	cache.Set("age", 18)
	cache.Set("data", map[string]interface{}{"a": 1})
	log.Println(cache.Get("name"))
	log.Println(cache.Get("age"))
	log.Println(cache.Get("data"))
	log.Println(cache.Get("int"))
	log.Println(cache.Del("int"))
	log.Println(cache.Del("age"))
	log.Println("Before Flush Keys:", cache.Keys())
	cache.Flush()
	log.Println("After Flush Keys:", cache.Keys())

	cache1 := cache_server.NewCacheServer()
	cache1.SetMaxMemory("100MB")
	cache1.Set("int", 2)
	log.Println(cache1.Get("int"))
}
