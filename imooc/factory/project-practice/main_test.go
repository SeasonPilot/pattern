package main

import (
	"fmt"
	"testing"
)

func TestMemCache(t *testing.T) {
	var cacheFactory CacheFactory

	// 给抽象工厂赋值为具体工厂, 方便切换底层不同缓存实现
	cacheFactory = &RedisFactory{} // RedisFactory
	//cacheFactory = &MemFactory{} // MemFactory

	cache := cacheFactory.Create()

	cache.Set("k1", "v1")
	fmt.Println(cache.Get("k1"))
}
