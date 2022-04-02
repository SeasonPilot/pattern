package main

import "errors"

// 定义一个cache接口，作为父类
type Cache interface {
	Set(k, v string)
	Get(string) string
}

// 实现具体的 cache: redis cache
type Redis struct {
	data map[string]string
}

func (r *Redis) Set(k string, v string) {
	r.data[k] = v
}

func (r *Redis) Get(k string) string {
	return "Redis data:" + r.data[k]
}

// 实现具体的 cache: MemCache cache
type MemCache struct {
	data map[string]string
}

func (m *MemCache) Set(k, v string) {
	m.data[k] = v
}

func (m *MemCache) Get(k string) string {
	return "Redis data:" + m.data[k]
}

// 实现 cache 的简单工厂
type CacheFactory struct{}

type CacheType int

const (
	redis CacheType = iota
	memcache
)

func (cf *CacheFactory) Create(cacheType CacheType) (Cache, error) {
	if cacheType == redis {
		return &Redis{data: map[string]string{}}, nil
	}
	if cacheType == memcache {
		return &MemCache{data: map[string]string{}}, nil
	}
	return nil, errors.New("error cache type")
}
