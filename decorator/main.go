package decorator

import "fmt"

// 抽象组件   相当于 http.Handler？
type Component interface {
	Operator()
}

// 具体组件  相当于 http.HandlerFunc？
type Component1 struct {
}

// 相当于 ServeHTTP？
func (c Component1) Operator() {
	fmt.Println("操作 1")
}

// 核心代码
// 抽象装饰器,要继承 被装饰的抽象类
type Decorator interface {
	Component // 要继承
	Do()      // 装饰器方法
}

// 具体装饰器,和被装饰的抽象类，因为下面 Operator 要调用被装饰的类
type Decorator1 struct {
	c Component // 这里是组合
}

// 装饰逻辑，只包含装饰相关的操作
func (d Decorator1) Do() {
	fmt.Println("装饰操作")
}

//
func (d Decorator1) Operator() {
	d.Do()         // 调用装饰器
	d.c.Operator() // 调用被装饰的类
}
