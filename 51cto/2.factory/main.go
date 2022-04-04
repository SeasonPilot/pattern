package main

import "fmt"

type Operator interface {
	SetLeft(int)
	SetRight(int)
	Result() int
}

// 抽象工厂
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	left, right int
}

func (o *OperatorBase) SetLeft(i int) {
	o.left = i
}

func (o *OperatorBase) SetRight(i int) {
	o.right = i
}

// 具体工厂
type PlusOperatorFactory struct{}

type PlusOperator struct {
	OperatorBase // 继承
}

func (o *PlusOperator) Result() int {
	return o.left + o.right
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{}
}

func main() {
	// 定义接口类型变量
	var fac OperatorFactory
	// 赋值
	fac = &PlusOperatorFactory{}

	operator := fac.Create()

	operator.SetLeft(20)
	operator.SetRight(10)
	fmt.Println(operator.Result())
}
