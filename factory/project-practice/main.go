package main

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

// 实现具体的 cache: MemFactory cache
type MemCache struct {
	data map[string]string
}

func (m *MemCache) Set(k, v string) {
	m.data[k] = v
}

func (m *MemCache) Get(k string) string {
	return "Mem data:" + m.data[k]
}

// 抽象工厂; 与简单工厂差异从这行开始；核心代码
type CacheFactory interface {
	Create() Cache
}

// 具体工厂函数
type RedisFactory struct{}

func (c *RedisFactory) Create() Cache {
	// Redis{} 里面有 map，要用 make 初始化
	return &Redis{
		data: make(map[string]string),
	}
}

type MemFactory struct{}

func (c *MemFactory) Create() Cache {
	return &MemCache{
		data: make(map[string]string),
	}
}
