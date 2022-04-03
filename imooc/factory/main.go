package main

type Product interface {
	SetName(string)
	GetName() string
}

type Product1 struct {
	Name string
}

// SetName 这里要改变 Product1 的值，所以是指针类型的接收者
func (p *Product1) SetName(s string) {
	p.Name = s
}

func (p *Product1) GetName() string {
	return "产品 1 的 Name 为：" + p.Name
}

type Product2 struct {
	Name string
}

func (p *Product2) SetName(s string) {
	p.Name = s
}

func (p *Product2) GetName() string {
	return "产品 2 的 Name 为：" + p.Name
}

// 抽象工厂
type Factory interface {
	Create() Product
}

// 实现一个具体的工厂: 产品 1 的工厂
type ProductFactory1 struct{}

func (pf *ProductFactory1) Create() Product {
	return &Product1{}
}

// 实现一个具体的工厂: 产品 2 的工厂
type ProductFactory2 struct{}

func (pf *ProductFactory2) Create() Product {
	return &Product2{}
}
