package template_pattern

import "fmt"

type Cooker interface {
	Fire()
	Cooke()
	outFire()
}

// BaseCooker 类似于一个抽象类
// 超类
type BaseCooker struct {
}

func (b BaseCooker) Fire() {
	fmt.Println("fire")
}

// Cooke 做菜，交给具体的子类实现
func (b BaseCooker) Cooke() {}

func (b BaseCooker) outFire() {
	fmt.Println("out fire")
}

// DoCook 封装具体步骤 fixme
func DoCook(c Cooker) {
	c.Fire()
	c.Cooke()
	c.outFire()
}

type MyCooker struct {
	BaseCooker
}

func (m MyCooker) Cooke() {
	fmt.Println("MyCooker")
}

type ChaoJiDan struct {
	BaseCooker
}

func (ChaoJiDan) Cooke() {
	fmt.Println("做炒鸡蛋")
}
