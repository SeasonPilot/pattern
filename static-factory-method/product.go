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

type productType int

const (
	product1 productType = iota
	product2
)

type ProductFactory struct {
}

// Create 入参要是自定义类型 productType
func (p ProductFactory) Create(s productType) Product {
	if s == product1 {
		return &Product1{}
	}
	if s == product2 {
		return &Product2{}
	}
	return nil
}
