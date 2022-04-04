package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManger struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManger() *PrototypeManger {
	return &PrototypeManger{make(map[string]Cloneable)}
}

func (p *PrototypeManger) Set(name string, prototype Cloneable) {
	p.prototypes[name] = prototype
}

func (p *PrototypeManger) Get(name string) Cloneable {
	return p.prototypes[name]
}

type Type1 struct {
	name string
}

// 深拷贝
func (p *Type1) Clone() Cloneable {
	tc := *p
	return &tc
}

// 浅拷贝
type Type2 struct {
	name string
}

func (p *Type2) Clone() Cloneable {
	return p
}

func main() {
	p := &Type2{
		name: "type2",
	}
	pm := NewPrototypeManger()
	pm.Set("t1", p)
	p1 := pm.Get("t1")
	p2 := p.Clone()

	if p1 == p2 {
		fmt.Println("浅拷贝")
	} else {
		fmt.Println("深拷贝")
	}
}
