package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

// 1. 在类中添加一个私有静态成员变量用于保存单例实例。
var singleInstance *single

// 2. 声明一个公有静态构建方法用于获取单例实例。
func getInstance() *single {
	// 3. 在静态方法中实现"延迟初始化"。 该方法会在首次被调用时创建一个新对象， 并将其存储在静态成员变量中。 此后该方法每次被调用时都返回该实例。
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				// 4. 将类的构造函数设为私有。 类的静态方法仍能调用构造函数， 但是其他对象不能调用。
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
