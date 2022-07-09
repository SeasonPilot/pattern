package options_pattern_practice2

import "time"

type Connect struct {
	Addr    string
	Cache   bool
	TimeOut time.Duration
}

type Options struct {
	Cache   bool
	TimeOut time.Duration
}

// Optioner overrides behavior of Connect.
type Optioner interface {
	apply(*Options)
}

type OptionFunc func(*Options)

func (f OptionFunc) apply(o *Options) {
	f(o)
}

// 确保 OptionFunc 类型 实现了接口
var _ Optioner = OptionFunc(func(options *Options) {})

func WithCache(cache bool) Optioner {
	return OptionFunc(func(options *Options) {
		options.Cache = cache
	})
}

func NewConn(addr string, opts ...Optioner) (Connect, error) {
	o := &Options{
		Cache:   false,
		TimeOut: 0,
	}

	for _, opt := range opts {
		opt.apply(o)
	}

	return Connect{
		Addr:    addr,
		Cache:   o.Cache,
		TimeOut: o.TimeOut,
	}, nil
}
