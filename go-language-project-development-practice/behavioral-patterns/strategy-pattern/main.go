package strategy_pattern

// IStrategy 定义接口
type IStrategy interface {
	Do(int, int) int
}

type Operator struct {
	strategy IStrategy // 组合接口
}

func (o *Operator) SetStrategy(strategy IStrategy) {
	o.strategy = strategy // 给接口赋值具体的实现类
}

func (o *Operator) Calculate(a, b int) int {
	return o.strategy.Do(a, b) // 调用实现类的方法
}

type Add struct{}

func (*Add) Do(a, b int) int {
	return a + b
}

type Reduce struct{}

func (*Reduce) Do(a, b int) int {
	return a - b
}
