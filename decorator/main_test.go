package decorator

import "testing"

func TestComponent1_Operator(t *testing.T) {
	// 初始化父类
	var c Component
	// 给父类附具体的值
	c = Component1{}
	c.Operator()
}
func TestDecorator1(t *testing.T) {
	// 组件
	c := Component1{}
	// 装饰器
	d := Decorator1{
		c: c,
	}
	d.Operator()
}
