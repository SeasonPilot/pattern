package proxy_pattern_practice

import "fmt"

type Seller interface {
	Sell()
}

type Station struct {
}

func (s Station) Sell() {
	fmt.Println("Station sell")
}

// StationProxy 基于组合原则， 也就是说一个对象应该将部分工作委派给另一个对象。
type StationProxy struct {
	station Station // fixme:  持有一个火车站对象
}

func (s StationProxy) Sell() {
	fmt.Println("StationProxy sell")
}
