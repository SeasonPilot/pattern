package main

import (
	"fmt"
	"testing"
)

func TestProduct1(t *testing.T) {
	// 抽象工厂
	var pf Factory

	// ProductFactory1 实现抽象工厂
	pf = &ProductFactory1{}

	// ProductFactory2 实现抽象工厂
	//pf = &ProductFactory2{}

	// ProductFactory1 工厂的对象 创建产品 1
	p := pf.Create()

	p.SetName("hahha")
	fmt.Println(p.GetName())
}
