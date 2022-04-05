package main

// 具体生成器 （Concrete Builders）
type IntBuilder struct {
	result int // 产品
}

func (b *IntBuilder) Part1() {
	b.result += 1
}

func (b *IntBuilder) Part2() {
	b.result += 2
}

func (b *IntBuilder) Part3() {
	b.result += 3

}

// 返回产品
func (b *IntBuilder) GetResult() int {
	return b.result
}
