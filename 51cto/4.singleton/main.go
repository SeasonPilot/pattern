package main

import (
	"fmt"
	"sync"
)

type Single struct {
	data int
}

var singleton *Single
var once sync.Once // 要放在全局变量。内核信号，时时刻刻只能运行一个

func GetInterface() *Single {
	once.Do(func() {
		singleton = &Single{data: 100}
	})
	return singleton
}

func main() {
	single1 := GetInterface()
	single2 := GetInterface()

	single2.data = 300
	fmt.Println(single1.data)

	if single1 == single2 {
		fmt.Println("实例地址相同")
	} else {
		fmt.Println("实例地址不相同")
	}
}
