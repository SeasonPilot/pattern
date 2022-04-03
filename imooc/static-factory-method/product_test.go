package main

import (
	"fmt"
	"testing"
)

func TestFactory_ProductFactory(t *testing.T) {
	pf := ProductFactory{}
	product := pf.Create(product1)
	product.SetName("p1")
	fmt.Println("结果：", product.GetName())
}

func TestProduct1_SetName(t *testing.T) {
	a := Product1{}
	a.SetName("product111")
	fmt.Println("结果：", a.GetName())
}
