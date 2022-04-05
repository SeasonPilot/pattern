package main

// 主管
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d Director) Make() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}
