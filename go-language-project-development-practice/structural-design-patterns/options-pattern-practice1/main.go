package options_pattern_practice1

import "time"

type Connection struct {
	Addr    string
	Cache   bool
	Timeout time.Duration
}

type Option struct {
	Cache   bool
	Timeout time.Duration
}

// 核心代码

type OptionFunc func(*Option)

func WithCache(cache bool) OptionFunc {
	return func(o *Option) {
		o.Cache = cache
	}
}

func NewConn(addr string, opts ...OptionFunc) Connection {
	o := &Option{ // fixme:这里不是 Connection
		Cache:   false,
		Timeout: 10,
	}

	for _, fun := range opts {
		fun(o)
	}

	return Connection{
		Addr:    addr,
		Cache:   o.Cache,
		Timeout: o.Timeout,
	}
}
