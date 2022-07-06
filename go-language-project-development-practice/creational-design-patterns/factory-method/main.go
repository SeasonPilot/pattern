package factory_method

type Person struct {
	Name string
	Age  int
}

// NewFactory 使⽤此功能来创建具有默认年龄的⼯⼚
func NewFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			Name: name,
			Age:  age,
		}
	}
}
